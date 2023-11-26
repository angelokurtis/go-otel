package otel

import (
	"context"
	"fmt"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/angelokurtis/go-otel/starter/internal/env"
	"github.com/angelokurtis/go-otel/starter/internal/logger"
	intltrace "github.com/angelokurtis/go-otel/starter/internal/trace"
)

// Providers struct holds the TracerProvider for OpenTelemetry.
type Providers struct {
	TracerProvider *sdktrace.TracerProvider
}

// SetupProviders initializes and configures OpenTelemetry providers.
// It returns a Providers struct containing the TracerProvider, a cleanup function, and an error if setup fails.
func SetupProviders(ctx context.Context) (*Providers, func(), error) {
	// Create a new OpenTelemetry resource.
	r, err := intltrace.NewResource(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	// Lookup environment variables related to OpenTelemetry configuration.
	variables, err := env.LookupVariables()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	// Create a sampler based on environment variables.
	s, err := intltrace.NewSampler(intltrace.SamplerOptions{
		Sampler:    env.ToTraceSampler(variables),
		SamplerArg: env.ToTraceSamplerArg(variables),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	// Create a client for exporting traces based on environment variables.
	c, err := intltrace.NewClient(intltrace.ClientOptions{
		Protocol:    env.ToTraceProtocol(variables),
		Endpoint:    env.ToTraceEndpoint(variables),
		Compression: env.ToTraceCompression(variables),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	// Create span exporters based on environment variables.
	se, err := intltrace.NewSpanExporters(ctx, intltrace.SpanExportersOptions{
		Exporters: env.ToTraceExporters(variables),
		Client:    c,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	// Create a text map propagator based on environment variables.
	p := intltrace.NewTextMapPropagator(env.ToTracePropagators(variables))
	log := logger.New()
	// Create a TracerProvider with the configured options.
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

	// Return the initialized TracerProvider and cleanup function.
	return provs, func() {
		cleanup()
	}, nil
}
