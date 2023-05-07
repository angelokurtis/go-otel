package config

import (
	"net/url"
	"time"

	"github.com/angelokurtis/go-starter-otel/internal/trace"
)

type Trace struct{ env *Env }

func NewTrace(env *Env) *Trace { return &Trace{env: env} }

func (t *Trace) TracesExporter() []trace.Exporter {
	return t.env.TracesExporter
}

func (t *Trace) ExporterOTLPTracesEndpoint() *url.URL {
	if t.env.ExporterOTLPTracesEndpoint != nil && t.env.ExporterOTLPTracesEndpoint.String() != "" {
		return t.env.ExporterOTLPTracesEndpoint
	}
	return t.env.ExporterOTLPEndpoint
}

func (t *Trace) ExporterOTLPTracesTimeout() time.Duration {
	if t.env.ExporterOTLPTracesTimeout != nil {
		return *t.env.ExporterOTLPTracesTimeout
	}
	return t.env.ExporterOTLPTimeout
}

func (t *Trace) ExporterOTLPTracesProtocol() trace.Protocol {
	if t.env.ExporterOTLPTracesProtocol != nil {
		return *t.env.ExporterOTLPTracesProtocol
	}
	return trace.Protocol(t.env.ExporterOTLPProtocol)
}

func (t *Trace) ExporterOTLPTracesCompression() trace.Compression {
	if t.env.ExporterOTLPTracesCompression != nil {
		return *t.env.ExporterOTLPTracesCompression
	}
	return trace.Compression(t.env.ExporterOTLPCompression)
}

func (t *Trace) TracesSampler() trace.Sampler {
	return t.env.TracesSampler
}

func (t *Trace) TracesSamplerArg() float64 {
	return t.env.TracesSamplerArg
}

func (t *Trace) Propagators() []trace.Propagator {
	return t.env.Propagators
}
