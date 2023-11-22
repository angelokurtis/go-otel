package main

import (
	"context"
	"log"
	"net/http"

	otel "github.com/angelokurtis/go-starter-otel"
)

func main() {
	// Initialize OpenTelemetry with environment variables
	_, shutdown, err := otel.SetupProviders(context.Background())
	if err != nil {
		log.Fatalf("Error initializing OpenTelemetry: %v", err)
	}
	defer shutdown()

	// Example HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx, span := otel.StartSpanFromContext(ctx)
		defer span.End()

		// Your application code here

		w.Write([]byte("Hello, World!"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
