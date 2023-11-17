//go:generate go-enum --nocase

package trace

// Compression defines the compression type to use on OTLP.
// ENUM(none, gzip)
type Compression string

// Exporter defines a trace exporter type responsible for delivering spans to external receivers.
// ENUM(zipkin, otlp, none, jaeger, logging)
type Exporter string

// Propagator determine which distributed tracing header formats are used, and which baggage propagation header formats are used.
// ENUM(tracecontext, baggage, b3, b3multi, jaeger, xray, ottrace)
type Propagator string

// Protocol defines the encoding of telemetry data and the protocol used to exchange spans between the client and the server.
// ENUM(grpc, http/protobuf)
type Protocol string

// Sampler configures whether spans will be recorded.
// ENUM(always_on, always_off, traceidratio, parentbased_always_on, parentbased_always_off, parentbased_traceidratio)
type Sampler string
