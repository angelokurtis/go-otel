//go:generate go-enum --nocase

package config

// Compression defines the compression type to use on OTLP.
// ENUM(none, gzip)
type Compression string

// Protocol defines the encoding of telemetry data and the protocol used to exchange data between the client and the server.
// ENUM(grpc, http/protobuf)
type Protocol string
