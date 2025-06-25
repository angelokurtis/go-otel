package env

import (
	"github.com/gotidy/ptr"

	"github.com/angelokurtis/go-otel/starter/internal/trace"
)

func ToTraceExporters(otel *Variables, provider ExporterProvider) trace.Exporters {
	if exporters, ok := provider.TraceExporters(); ok {
		return exporters
	}

	return otel.Traces.Exporter
}

func ToTraceEndpoint(otel *Variables) trace.Endpoint {
	ep := otel.Exporter.OTLP.Traces.Endpoint
	def := otel.Exporter.OTLP.Endpoint

	return trace.Endpoint(ptr.ToDef(ep, def))
}

func ToTraceTimeout(otel *Variables) trace.BatchTimeout {
	t := otel.Exporter.OTLP.Traces.Timeout
	def := otel.Exporter.OTLP.Timeout

	return trace.BatchTimeout(ptr.ToDef(t, def))
}

func ToTraceProtocol(otel *Variables) trace.Protocol {
	p := otel.Exporter.OTLP.Traces.Protocol
	def := otel.Exporter.OTLP.Protocol

	return ptr.ToDef(p, trace.Protocol(def))
}

func ToTraceCompression(otel *Variables) trace.Compression {
	c := otel.Exporter.OTLP.Traces.Compression
	def := trace.Compression(otel.Exporter.OTLP.Compression)

	return ptr.ToDef(c, def)
}

func ToTraceSampler(otel *Variables) trace.Sampler {
	return otel.Traces.Sampler.Type
}

func ToTraceSamplerArg(otel *Variables) trace.SamplerArg {
	return trace.SamplerArg(otel.Traces.Sampler.Arg)
}

func ToTracePropagators(otel *Variables) trace.Propagators {
	return otel.Propagators
}
