package otel

import (
	"context"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/angelokurtis/go-starter-otel/internal/env"
	"github.com/angelokurtis/go-starter-otel/internal/trace"
)

type Provider struct {
	Tracer *sdktrace.TracerProvider

	propagation propagation.TextMapPropagator
}

func Init(ctx context.Context) (*Provider, func(), error) {
	r := resource.Environment()

	o, err := env.LookupOTel()
	if err != nil {
		return nil, nil, err
	}

	config := env.ToTrace(o)

	s, err := trace.NewSampler(config)
	if err != nil {
		return nil, nil, err
	}

	c, err := trace.NewTraceClient(config)
	if err != nil {
		return nil, nil, err
	}

	se, err := trace.NewSpanExporters(ctx, config, c)
	if err != nil {
		return nil, nil, err
	}

	tp, cleanup := trace.NewTracerProvider(ctx, r, s, se, config)
	p := trace.NewTextMapPropagator(config)
	provider := &Provider{
		Tracer:      tp,
		propagation: p,
	}

	return provider, func() {
		cleanup()
	}, nil
}
