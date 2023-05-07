//go:generate go-enum --nocase

package metric

// Compression defines the compression type to use on OTLP.
// ENUM(none, gzip)
type Compression string

// Exporter defines a metric exporter type responsible for delivering metric data to external receivers.
// ENUM(otlp, none, prometheus, logging)
type Exporter string

// Protocol defines the encoding of telemetry data and the protocol used to exchange metric data between the client and the server.
// ENUM(g rpc, http/protobuf)
type Protocol string
