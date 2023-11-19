package otel

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/angelokurtis/go-starter-otel/internal/env"
	"github.com/angelokurtis/go-starter-otel/internal/logger"
	intltrace "github.com/angelokurtis/go-starter-otel/internal/trace"
)

type Providers struct {
	TracerProvider *sdktrace.TracerProvider

	propagation propagation.TextMapPropagator
}

func SetupProviders(ctx context.Context) (*Providers, func(), error) {
	r, err := intltrace.NewResource(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	o, err := env.LookupOTel()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	config := env.ToTrace(o)

	s, err := intltrace.NewSampler(config)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	c, err := intltrace.NewTraceClient(config)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	se, err := intltrace.NewSpanExporters(ctx, config, c)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	l := logger.New()
	tp, cleanup := intltrace.NewTracerProvider(ctx, l, r, s, se, config)
	p := intltrace.NewTextMapPropagator(config)
	provider := &Providers{
		TracerProvider: tp,
		propagation:    p,
	}

	return provider, func() {
		cleanup()
	}, nil
}
