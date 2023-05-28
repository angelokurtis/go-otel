package otel

import (
	"context"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// Provider defines an interface for shutting down an OpenTelemetry provider.
type Provider interface {
	// Shutdown gracefully shuts down the OpenTelemetry provider.
	Shutdown(ctx context.Context)
}

type provider struct {
	tracer *sdktrace.TracerProvider
}

func (p *provider) Shutdown(ctx context.Context) {
	_ = p.tracer.Shutdown(ctx)
}
