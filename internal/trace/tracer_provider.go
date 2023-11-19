package trace

import (
	"context"

	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewTracerProvider(ctx context.Context, logger logr.Logger, res *resource.Resource, sampler trace.Sampler, exporters []trace.SpanExporter, config *Config) (*trace.TracerProvider, func()) {
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
		if err := tp.Shutdown(ctx); err != nil {
			logger.V(1).Info("Error shutting down Tracer Provider", "err", err)
		}
	}

	return tp, cleanup
}
