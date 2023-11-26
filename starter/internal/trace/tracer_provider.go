package trace

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

type TracerProviderOptions struct {
	Resource     *resource.Resource
	Sampler      trace.Sampler
	Exporters    []trace.SpanExporter
	BatchTimeout BatchTimeout
	Propagator   propagation.TextMapPropagator
	Logger       logr.Logger
}

func NewTracerProvider(ctx context.Context, options TracerProviderOptions) (*trace.TracerProvider, func()) {
	opts := []trace.TracerProviderOption{
		trace.WithResource(options.Resource),
		trace.WithSampler(options.Sampler),
	}
	for _, exporter := range options.Exporters {
		opts = append(opts, trace.WithBatcher(
			exporter,
			trace.WithBatchTimeout(time.Duration(options.BatchTimeout))),
		)
	}

	tp := trace.NewTracerProvider(opts...)

	otel.SetTextMapPropagator(options.Propagator)
	otel.SetLogger(options.Logger)
	otel.SetTracerProvider(tp)

	cleanup := func() {
		if err := tp.Shutdown(ctx); err != nil {
			options.Logger.V(1).Info("Error shutting down Tracer Provider", "err", err.Error())
		}
	}

	return tp, cleanup
}
