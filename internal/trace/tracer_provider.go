package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewTracerProvider(ctx context.Context, res *resource.Resource, sampler trace.Sampler, exporters []trace.SpanExporter, config *Config) (*trace.TracerProvider, func()) {
	options := []trace.TracerProviderOption{
		trace.WithResource(res),
		trace.WithSampler(sampler),
	}
	for _, exp := range exporters {
		options = append(options, trace.WithBatcher(
			exp,
			trace.WithBatchTimeout(config.Timeout)),
		)
	}

	tp := trace.NewTracerProvider(options...)
	otel.SetTracerProvider(tp)

	cleanup := func() {
		_ = tp.Shutdown(ctx)
	}

	return tp, cleanup
}
