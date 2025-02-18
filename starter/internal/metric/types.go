//go:generate go-enum --nocase

package metric

import (
	"net/url"
)

// Exporters is a slice of Exporter, representing a collection of metric exporters.
type Exporters []Exporter

// Endpoint represents a URL endpoint for metric data export.
type Endpoint url.URL

// Compression defines the compression type to use on OTLP.
// ENUM(none, gzip)
type Compression string

// Exporter defines a metric exporter type responsible for delivering metric data to external receivers.
// ENUM(otlp, none, prometheus, logging)
type Exporter string

// Protocol defines the encoding of telemetry data and the protocol used to exchange metric data between the client and the server.
// ENUM(grpc, http/protobuf)
type Protocol string

// PrometheusHost defines the hostname or IP address where metrics are published.
type PrometheusHost string

// PrometheusPort specifies the port on which metrics are published.
type PrometheusPort uint

// PrometheusPath specifies the URL path where Prometheus scrapes the published metrics from.
type PrometheusPath string
