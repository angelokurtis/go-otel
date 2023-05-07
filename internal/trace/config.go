package trace

import (
	"net/url"
	"time"
)

type Config interface {
	TracesExporter() []Exporter
	ExporterOTLPTracesEndpoint() *url.URL
	ExporterOTLPTracesTimeout() time.Duration
	ExporterOTLPTracesProtocol() Protocol
	ExporterOTLPTracesCompression() Compression
	TracesSampler() Sampler
	TracesSamplerArg() float64
	Propagators() []Propagator // TODO: setup propagators
}
