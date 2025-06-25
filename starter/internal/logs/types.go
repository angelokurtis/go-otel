//go:generate go-enum --nocase

package logs

import (
	"net/url"
	"time"
)

// Exporters is a slice of Exporter instances representing multiple log exporters.
type Exporters []Exporter

// Endpoint represents a log endpoint as a URL.
type Endpoint url.URL

// Timeout defines the maximum waiting time allowed to send each OTLP log batch.
type Timeout time.Duration

// Compression defines the compression type to use on OTLP.
// ENUM(none, gzip)
type Compression string

// Exporter defines a log exporter type responsible for delivering logs to external receivers.
// ENUM(otlp, none, console)
type Exporter string

// Protocol defines the encoding of telemetry data and the protocol used to exchange logs between the client and the server.
// ENUM(grpc, http/protobuf)
type Protocol string

// ExportInterval defines the delay interval between two consecutive log exports.
type ExportInterval time.Duration

// ExportTimeout defines the maximum allowed time to export logs.
type ExportTimeout time.Duration

// MaxQueueSize defines the maximum number of logs that can be queued for processing.
type MaxQueueSize int

// ExportMaxBatchSize defines the maximum number of logs in a single export batch.
type ExportMaxBatchSize int
