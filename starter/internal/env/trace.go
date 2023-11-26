package env

import (
	"github.com/gotidy/ptr"

	"github.com/angelokurtis/go-otel/starter/internal/trace"
)

func ToTraceExporters(otel *Variables) trace.Exporters {
	return otel.Traces.Exporter
}

func ToTraceEndpoint(otel *Variables) trace.Endpoint {
	return trace.Endpoint(ptr.ToDef(otel.Exporter.OTLP.Traces.Endpoint, otel.Exporter.OTLP.Endpoint))
}

func ToTraceTimeout(otel *Variables) trace.BatchTimeout {
	return trace.BatchTimeout(ptr.ToDef(otel.Exporter.OTLP.Traces.Timeout, otel.Exporter.OTLP.Timeout))
}

func ToTraceProtocol(otel *Variables) trace.Protocol {
	return ptr.ToDef(otel.Exporter.OTLP.Traces.Protocol, trace.Protocol(otel.Exporter.OTLP.Protocol))
}

func ToTraceCompression(otel *Variables) trace.Compression {
	return ptr.ToDef(otel.Exporter.OTLP.Traces.Compression, trace.Compression(otel.Exporter.OTLP.Compression))
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
