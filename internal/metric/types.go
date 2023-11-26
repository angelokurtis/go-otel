//go:generate go-enum --nocase

package metric

import (
	"net/url"
	"time"
)

// Exporters is a slice of Exporter, representing a collection of metric exporters.
type Exporters []Exporter

// Endpoint represents a URL endpoint for metric data export.
type Endpoint url.URL

// ExportInterval represents the interval between metric exports.
type ExportInterval time.Duration

// Compression defines the compression type to use on OTLP.
// ENUM(none, gzip)
type Compression string

// Exporter defines a metric exporter type responsible for delivering metric data to external receivers.
// ENUM(otlp, none, prometheus, logging)
type Exporter string

// Protocol defines the encoding of telemetry data and the protocol used to exchange metric data between the client and the server.
// ENUM(grpc, http/protobuf)
type Protocol string
