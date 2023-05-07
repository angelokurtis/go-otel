package otel

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/angelokurtis/go-starter-otel/internal/config"
	"github.com/angelokurtis/go-starter-otel/internal/trace"
)

// Provider defines an interface for shutting down an OpenTelemetry provider.
type Provider interface {
	// Shutdown gracefully shuts down the OpenTelemetry provider.
	Shutdown(ctx context.Context)
}

// Init is an initializer for an OpenTelemetry provider that offers a high level of configurability
func Init(ctx context.Context) (Provider, error) {
	res := resource.Environment()

	env, err := config.NewFromEnv()
	if err != nil {
		return nil, err
	}

	traceConfig := config.NewTrace(env)

	sampler, err := trace.NewSampler(ctx, traceConfig)
	if err != nil {
		return nil, err
	}

	traceClient, err := trace.NewTraceClient(traceConfig)
	if err != nil {
		return nil, err
	}

	spanExp, err := trace.NewSpanExporters(ctx, traceConfig, traceClient)
	if err != nil {
		return nil, err
	}

	tracerProvider := trace.NewTracerProvider(res, sampler, spanExp, traceConfig)

	return &provider{tracer: tracerProvider}, nil
}

type provider struct {
	tracer *sdktrace.TracerProvider
}

func (p *provider) Shutdown(ctx context.Context) {
	_ = p.tracer.Shutdown(ctx)
}
