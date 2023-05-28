package env

import (
	"github.com/gotidy/ptr"

	"github.com/angelokurtis/go-starter-otel/internal/trace"
)

func ToTrace(otel *OTel) *trace.Config {
	return &trace.Config{
		Exporters:   otel.Traces.Exporter,
		Endpoint:    ptr.ToDef(otel.Exporter.OTLP.Traces.Endpoint, otel.Exporter.OTLP.Endpoint),
		Timeout:     ptr.ToDef(otel.Exporter.OTLP.Traces.Timeout, otel.Exporter.OTLP.Timeout),
		Protocol:    ptr.ToDef(otel.Exporter.OTLP.Traces.Protocol, trace.Protocol(otel.Exporter.OTLP.Protocol)),
		Compression: ptr.ToDef(otel.Exporter.OTLP.Traces.Compression, trace.Compression(otel.Exporter.OTLP.Compression)),
		Sampler:     otel.Traces.Sampler.Type,
		SamplerArg:  otel.Traces.Sampler.Arg,
		Propagators: otel.Propagators,
	}
}
