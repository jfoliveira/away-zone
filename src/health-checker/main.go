package main

import (
	"context"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// We need an execution context to keep health-checker running and be able to shutdown gracefully.
	// Here using NotifyContext for convenience, as it resets the signal and mark the context as done
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := log.New(os.Stdout, "- HCS - ", log.Lmsgprefix|log.LstdFlags)
	logger.Printf("INFO: Starting health checker ...")

	// We need a single HTTP client to share common HTTP requests configuration values used
	// in all health check calls
	client := &http.Client{}

	// TODO: get list of URLs from env vars
	var urls = []string{
		"http://127.0.0.1:8000?param1",
		"http://127.0.0.1:8000?param2",
		"http://127.0.0.1:8000?param=value",
		"http://127.0.0.1:8000?param1=value2",
	}

	// TO DO: run in configurable intervals
	// We need a wait group to keep track of running Go routines
	var wg sync.WaitGroup

	wg.Add(len(urls))
	for _, url := range urls {
		go func(currentUrl string) {
			defer wg.Done()
			go checkEndpoint("GET", currentUrl, client, logger)
		}(url)
	}
	logger.Print("INFO: Ready! Waiting for check results ...")

	wg.Wait()

	<-ctx.Done()
	shutdown(ctx, logger)

	logger.Print("INFO: Bye bye! TTFN! :-)")
}

func checkEndpoint(method string, endpointUrl string, client *http.Client, logger *log.Logger) {
	logger.Printf("INFO: Starting checkpoint for %s", endpointUrl)
	time.Sleep(time.Duration(rand.IntN(5)) * time.Second)
	request, err := http.NewRequest(method, endpointUrl, nil)
	if err != nil {
		logger.Printf("Error preparing request. Error: %s", err)
	}
	response, err := client.Do(request)

	if err != nil {
		logger.Printf("Error processing request to endpoint %s. Error: %s", request.URL, err)
	} else {
		defer response.Body.Close()

		if response.StatusCode != 200 {
			logger.Printf(
				"ERROR: Failed health check for endpoint %s. Expected HTTP status code: 200, Got: %d",
				request.RequestURI,
				response.StatusCode,
			)
		}
		logger.Printf(
			"INFO: Response: %d - %s %s",
			response.StatusCode,
			request.Method,
			request.URL,
		)

		for i := 0; i < 5; i++ {
			logger.Printf(
				"INFO: Routine iteration %d for %s",
				i,
				request.URL,
			)
			time.Sleep(time.Second)
		}
	}
}

// Do something on shutting down
func shutdown(ctx context.Context, logger *log.Logger) {
	logger.Printf("Shutting down ... Reason: %s", ctx.Err().Error())
}
