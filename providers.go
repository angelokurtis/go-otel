package otel

import (
	"context"
	"fmt"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/angelokurtis/go-starter-otel/internal/env"
	"github.com/angelokurtis/go-starter-otel/internal/logger"
	intltrace "github.com/angelokurtis/go-starter-otel/internal/trace"
)

type Providers struct {
	TracerProvider *sdktrace.TracerProvider
}

func SetupProviders(ctx context.Context) (*Providers, func(), error) {
	r, err := intltrace.NewResource(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	variables, err := env.LookupVariables()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	s, err := intltrace.NewSampler(intltrace.SamplerOptions{
		Sampler:    env.ToTraceSampler(variables),
		SamplerArg: env.ToTraceSamplerArg(variables),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	c, err := intltrace.NewClient(intltrace.ClientOptions{
		Protocol:    env.ToTraceProtocol(variables),
		Endpoint:    env.ToTraceEndpoint(variables),
		Compression: env.ToTraceCompression(variables),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	se, err := intltrace.NewSpanExporters(ctx, intltrace.SpanExportersOptions{
		Exporters: env.ToTraceExporters(variables),
		Client:    c,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	p := intltrace.NewTextMapPropagator(env.ToTracePropagators(variables))
	log := logger.New()
	tp, cleanup := intltrace.NewTracerProvider(ctx, intltrace.TracerProviderOptions{
		Resource:     r,
		Sampler:      s,
		Exporters:    se,
		BatchTimeout: env.ToTraceTimeout(variables),
		Propagator:   p,
		Logger:       log,
	})
	provs := &Providers{
		TracerProvider: tp,
	}

	return provs, func() {
		cleanup()
	}, nil
}
