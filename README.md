# go-otel

`go-otel` is a collection of libraries designed to facilitate the integration of OpenTelemetry into Go applications.
This project addresses the need for streamlined observability tools in Go, offering functionalities similar to those
available for auto-instrumentation in other languages. Our goal is to minimize setup complexity and provide easy access
to advanced observability features with minimal configuration.

## Libraries Overview

- **[starter](starter)**: Simplifies the configuration of OpenTelemetry tracing and metrics through environment
  variables, enabling nearly automatic instrumentation.
- **[span](span)**: Offers a simplified interface for managing span operations, enhancing the ease of logging and
  manipulating span data.

## Quick Start

### Requirements

- Go version 1.23 or newer
- Docker (optional, for running specific examples or tests)

### Basic Usage

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

Set the necessary environment variables listed below, then run your Go application. OpenTelemetry will configure itself
automatically.

## Configuration Options

| Environment Variable                     | Description                                                           | Default                 |
|------------------------------------------|-----------------------------------------------------------------------|-------------------------|
| `OTEL_TRACES_EXPORTER`                   | Exporter for traces (`otlp`, `jaeger`, `zipkin`, `logging`)           | `otlp`                  |
| `OTEL_METRICS_EXPORTER`                  | Exporter for metrics (`none`, `otlp`, `prometheus`)                   | `none`                  |
| `OTEL_LOGS_EXPORTER`                     | Exporter for logs (`none`, `otlp`)                                    | `none`                  |
| `OTEL_EXPORTER_OTLP_ENDPOINT`            | Unified endpoint for all signals                                      | `http://localhost:4317` |
| `OTEL_EXPORTER_OTLP_PROTOCOL`            | Transport protocol (`grpc`, `http/protobuf`)                          | `grpc`                  |
| `OTEL_EXPORTER_OTLP_TIMEOUT`             | Timeout for all signal exports                                        | `10s`                   |
| `OTEL_EXPORTER_OTLP_TRACES_ENDPOINT`     | Endpoint for trace requests                                           | `http://localhost:4317` |
| `OTEL_EXPORTER_OTLP_TRACES_PROTOCOL`     | Protocol for trace requests                                           | `grpc`                  |
| `OTEL_EXPORTER_OTLP_TRACES_TIMEOUT`      | Timeout for trace requests                                            | `10s`                   |
| `OTEL_EXPORTER_OTLP_METRICS_ENDPOINT`    | Endpoint for metric requests                                          | `http://localhost:4317` |
| `OTEL_EXPORTER_OTLP_METRICS_PROTOCOL`    | Protocol for metric requests                                          | `grpc`                  |
| `OTEL_EXPORTER_OTLP_METRICS_TIMEOUT`     | Timeout for metric requests                                           | `10s`                   |
| `OTEL_EXPORTER_OTLP_LOGS_ENDPOINT`       | Endpoint for log requests                                             | `http://localhost:4317` |
| `OTEL_EXPORTER_OTLP_LOGS_PROTOCOL`       | Protocol for log requests                                             | `grpc`                  |
| `OTEL_EXPORTER_OTLP_LOGS_TIMEOUT`        | Timeout for log requests                                              | `10s`                   |
| `OTEL_EXPORTER_PROMETHEUS_HOST`          | Host to bind Prometheus server                                        | `0.0.0.0`               |
| `OTEL_EXPORTER_PROMETHEUS_PORT`          | Port for Prometheus server                                            | `9464`                  |
| `OTEL_EXPORTER_PROMETHEUS_PATH`          | Prometheus metrics path                                               | `/metrics`              |
| `OTEL_BLRP_EXPORT_TIMEOUT`               | Max time to export data                                               | `30s`                   |
| `OTEL_BLRP_SCHEDULE_DELAY`               | Delay between exports                                                 | `5s`                    |
| `OTEL_BLRP_MAX_QUEUE_SIZE`               | Max queued spans                                                      | `2048`                  |
| `OTEL_BLRP_MAX_EXPORT_BATCH_SIZE`        | Max spans in a batch                                                  | `512`                   |
| `OTEL_EXPORTER_OTLP_TRACES_COMPRESSION`  | Compression for trace requests                                        | `gzip`                  |
| `OTEL_EXPORTER_OTLP_METRICS_COMPRESSION` | Compression for metric requests                                       | `gzip`                  |
| `OTEL_EXPORTER_OTLP_LOGS_COMPRESSION`    | Compression for log requests                                          | `gzip`                  |
| `OTEL_EXPORTER_OTLP_COMPRESSION`         | Compression type (`gzip`, etc)                                        | `gzip`                  |
| `OTEL_PROPAGATORS`                       | Comma-separated list of propagators (`tracecontext`, `baggage`, `b3`) | `tracecontext,baggage`  |
| `OTEL_TRACES_SAMPLER`                    | Sampler type (`always_on`, `parentbased_traceidratio`)                | `parentbased_always_on` |
| `OTEL_TRACES_SAMPLER_ARG`                | Sampler argument (e.g., `0.25` for 25% sampling)                      | -                       |

## Contributing

We welcome contributions! Please feel free to submit issues or pull requests on
our [GitHub repository](https://github.com/angelokurtis/go-otel).

## License

This project is licensed under the Apache License 2.0. For details, see the [LICENSE](LICENSE) file.
