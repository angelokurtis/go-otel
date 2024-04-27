# go-otel

`go-otel` is a collection of libraries designed to facilitate the integration of OpenTelemetry into Go applications.
This project addresses the need for streamlined observability tools in Go, offering functionalities similar to those
available for auto-instrumentation in other languages, such as Java. Our goal is to minimize setup complexity and
provide easy access to advanced observability features with minimal configuration.

## Libraries Overview

- **[starter](starter)**: Simplifies the configuration of OpenTelemetry tracing and metrics through environment
  variables, enabling nearly automatic instrumentation.
- **[span](span)**: Offers a simplified interface for managing span operations, enhancing the ease of logging and
  manipulating span data.

## Quick Start

### Requirements

- Go version 1.21 or newer
- Docker (optional, for running specific examples or tests)

### Basic Usage

1. Initialize OpenTelemetry using environment variable-based configuration:

    ```go
    package main
    
    import (
        "context"
        "log"
        "net/http"
    
        "github.com/angelokurtis/go-otel/span"
        "github.com/angelokurtis/go-otel/starter"
    )
    
    func main() {
        // Initialize OpenTelemetry based on environment variables
        _, shutdown, err := starter.StartProviders(context.Background())
        if err != nil {
            log.Fatalf("Failed to initialize OpenTelemetry: %v", err)
        }
        defer shutdown()
    
        // Example HTTP server
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
            ctx := r.Context()
            ctx, end := span.Start(ctx)
            defer end()
    
            // Place your application code here

            w.Write([]byte("Hello, World!"))
        })
    
        log.Fatal(http.ListenAndServe(":8080", nil))
    }
    ```

2. Set the necessary environment variables as detailed in
   the [OpenTelemetry Autoconfigure Java SDK README](https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md).

3. Execute your Go application. OpenTelemetry will configure itself based on the specified environment variables.

Explore the `/_examples` directory for detailed guides on using each library.

## Configuration Options

Configuration aligns with the environment variables specified in
the [OpenTelemetry Autoconfigure Java SDK README](https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md).
Refer to this document for a detailed list of configuration options.

## Contributing

We welcome contributions! Please feel free to submit issues or pull requests on
our [GitHub repository](https://github.com/angelokurtis/go-otel).

## License

This project is licensed under the Apache License 2.0. For details, see the [LICENSE](LICENSE) file.
