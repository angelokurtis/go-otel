# go-starter-otel

**go-starter-otel** is a Go library that simplifies the configuration of OpenTelemetry by relying on environment
variables. It draws inspiration from the OpenTelemetry Autoconfigure Java SDK and allows users to configure various
aspects of OpenTelemetry instrumentation seamlessly using environment variables.

## Overview

OpenTelemetry is an observability framework that provides APIs, libraries, agents, and instrumentation to provide
observability for applications. go-starter-otel aims to streamline the setup process by allowing users to configure
OpenTelemetry using environment variables. This is particularly useful for projects where configuration through code or
configuration files may be cumbersome or less flexible.

## Features

- **Environment Variable Based Configuration:** Configure OpenTelemetry settings using environment variables, following
  the conventions defined in the OpenTelemetry Autoconfigure Java SDK.

- **Seamless Integration:** Easily integrate OpenTelemetry into your Go applications without the need for extensive code
  modifications.

## Getting Started

### Prerequisites

Make sure you have Go installed on your system.

### Installation

To install go-starter-otel, use the following Go module command:

```bash
go get github.com/angelokurtis/go-starter-otel
```

### Usage

1. Import the `go-starter-otel` package into your Go code:

    ```go
    import "github.com/angelokurtis/go-starter-otel"
    ```

2. Initialize OpenTelemetry with environment variable-based configuration:

    ```go
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
    ```

3. Set the required environment variables based on the configuration options described in
   the [OpenTelemetry Autoconfigure Java SDK README](https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md).

4. Run your Go application, and OpenTelemetry will be configured based on the provided environment variables.

## Configuration Options

The configuration options for go-starter-otel are aligned with the environment variables specified in
the [OpenTelemetry Autoconfigure Java SDK README](https://github.com/open-telemetry/opentelemetry-java/blob/main/sdk-extensions/autoconfigure/README.md).
Refer to that documentation for a comprehensive list of available options.

## Contributing

Contributions are welcome! Feel free to open issues or pull requests on
the [GitHub repository](https://github.com/angelokurtis/go-starter-otel).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.