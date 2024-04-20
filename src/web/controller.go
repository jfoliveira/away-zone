package main

import (
	"fmt"
	"html"
	"net/http"
)

// controller implements one method for each HTTP endpoint exposed by this mini web application
type controller struct {
}

func (c *controller) health() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w, "Healthy?: Endpoint %s\n", html.EscapeString(request.URL.Path))
	})
}
func (c *controller) index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w, "Welcome to Jef's minimal web application. This is the response for endpoint: %s\n", html.EscapeString(request.URL.Path))
	})
}
func (c *controller) status() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w, "Status: Endpoint %s\n", html.EscapeString(request.URL.Path))
	})
}
