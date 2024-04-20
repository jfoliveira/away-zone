package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type requestIDKeyValue string

const (
	requestIDKey requestIDKeyValue = ""
)

func main() {
	// TO DO: port should not be hardcoded
	listeningPort := 8000

	logger := log.New(os.Stdout, "", log.LstdFlags)
	logger.Printf("Starting HTTP server ... Listening on port %d", listeningPort)

	c := new(controller)
	router := http.NewServeMux()
	router.Handle("/", c.index())
	router.Handle("/health", c.health())
	router.Handle("/status", c.status())

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", listeningPort),
		Handler:      setUniqueRequestIDForTracing()(logIncomingRequest(logger)(router)),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on port %d: %v\n", listeningPort, err)
	}
}

// setUniqueRequestIDForTracing ensures every request is assigned a unique ID
func setUniqueRequestIDForTracing() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = getNextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// getNextRequestID generates a unique string to be used as an identifier for incoming requests
func getNextRequestID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// logIncomingRequest writes details of an incoming request using the given `logger`
func logIncomingRequest(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

