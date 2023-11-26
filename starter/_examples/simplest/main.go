package main

import (
	"context"
	"log"
	"net/http"

	otelstarter "github.com/angelokurtis/go-otel/starter"
)

func main() {
	// Initialize OpenTelemetry with environment variables
	_, shutdown, err := otelstarter.SetupProviders(context.Background())
	if err != nil {
		log.Fatalf("Error initializing OpenTelemetry: %v", err)
	}
	defer shutdown()

	// Example HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Your application code here

		w.Write([]byte("Hello, World!"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
