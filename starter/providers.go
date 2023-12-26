package starter

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/angelokurtis/go-otel/starter/internal/env"
	"github.com/angelokurtis/go-otel/starter/internal/logger"
	intlmetric "github.com/angelokurtis/go-otel/starter/internal/metric"
	intltrace "github.com/angelokurtis/go-otel/starter/internal/trace"
)

// Providers struct holds the TracerProvider for OpenTelemetry.
type Providers struct {
	TracerProvider *trace.TracerProvider
	MeterProvider  *metric.MeterProvider
}

type ShutdownFunc func()

// StartProviders initializes and configures OpenTelemetry providers.
// It returns a Providers struct containing the TracerProvider, a shutdown function, and an error if setup fails.
func StartProviders(ctx context.Context) (*Providers, ShutdownFunc, error) {
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

	// Create a logger for tracing and metrics.
	log := logger.New()

	// Create a TracerProvider with configured options.
	p := intltrace.NewTextMapPropagator(env.ToTracePropagators(variables))
	tracerProvider, traceCleanup := intltrace.NewTracerProvider(ctx, intltrace.TracerProviderOptions{
		Resource:     r,
		Sampler:      s,
		Exporters:    se,
		BatchTimeout: env.ToTraceTimeout(variables),
		Propagator:   p,
		Logger:       log,
	})

	// Create metric readers based on environment variables.
	readers, err := intlmetric.NewReaders(ctx, intlmetric.ReadersOptions{
		Exporters:   env.ToMetricExporters(variables),
		Endpoint:    env.ToMetricEndpoint(variables),
		Compression: env.ToMetricCompression(variables),
		Protocol:    env.ToMetricProtocol(variables),
	})
	if err != nil {
		traceCleanup()
		return nil, nil, fmt.Errorf("failed to set up OpenTelemetry providers: %w", err)
	}

	// Create a MeterProvider with configured options.
	meterProvider, metricCleanup := intlmetric.NewMeterProvider(ctx, intlmetric.MeterProviderOptions{
		Resource: r,
		Readers:  readers,
		Logger:   log,
	})
	provs := &Providers{
		TracerProvider: tracerProvider,
		MeterProvider:  meterProvider,
	}

	// Return the configured providers, cleanup functions, and no error.
	return provs, func() {
		metricCleanup()
		traceCleanup()
	}, nil
}
