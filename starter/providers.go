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

// Providers holds tracing and metrics providers for OpenTelemetry.
type Providers struct {
	TracerProvider *trace.TracerProvider
	MeterProvider  *metric.MeterProvider
}

type ShutdownFunc func()

// StartProviders initializes and configures OpenTelemetry providers.
// It returns a Providers struct containing the TracerProvider, a shutdown function, and an error if setup fails.
func StartProviders(ctx context.Context, opts ...func(c *option)) (*Providers, ShutdownFunc, error) {
	// Apply any configuration options
	options := newOption(opts...)

	// Create a resource to label all telemetry data
	resource, err := intltrace.NewResource(ctx, options)
	if err != nil {
		return nil, nil, fmt.Errorf("OpenTelemetry setup failed: %w", err)
	}

	// Load environment variables for configuration
	variables, err := env.LookupVariables()
	if err != nil {
		return nil, nil, fmt.Errorf("OpenTelemetry setup failed: %w", err)
	}

	// Set up a trace sampler (controls what traces are captured)
	sampler, err := intltrace.NewSampler(intltrace.SamplerOptions{
		Sampler:    env.ToTraceSampler(variables),
		SamplerArg: env.ToTraceSamplerArg(variables),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("OpenTelemetry setup failed: %w", err)
	}

	// Set up a trace client for exporting traces
	traceClient, err := intltrace.NewClient(intltrace.ClientOptions{
		Protocol:    env.ToTraceProtocol(variables),
		Endpoint:    env.ToTraceEndpoint(variables),
		Compression: env.ToTraceCompression(variables),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("OpenTelemetry setup failed: %w", err)
	}

	// Configure trace exporters for sending traces out
	spanExporters, err := intltrace.NewSpanExporters(ctx, intltrace.SpanExportersOptions{
		Exporters: env.ToTraceExporters(variables, options),
		Client:    traceClient,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("OpenTelemetry setup failed: %w", err)
	}

	// Initialize a logger for tracing and metrics
	traceLogger := logger.New()

	// Set up the TracerProvider to manage traces
	propagator := intltrace.NewTextMapPropagator(env.ToTracePropagators(variables))
	tracerProvider, tracerCleanup := intltrace.NewTracerProvider(ctx, intltrace.TracerProviderOptions{
		Resource:     resource,
		Sampler:      sampler,
		Exporters:    spanExporters,
		BatchTimeout: env.ToTraceTimeout(variables),
		Propagator:   propagator,
		Logger:       traceLogger,
	})

	// Configure metric readers for gathering metrics
	metricReaders, readersCleanup, err := intlmetric.NewReaders(ctx, intlmetric.ReadersOptions{
		Exporters:        env.ToMetricExporters(variables, options),
		Endpoint:         env.ToMetricEndpoint(variables),
		Compression:      env.ToMetricCompression(variables),
		Protocol:         env.ToMetricProtocol(variables),
		RegistryProvider: options,
		PrometheusHost:   env.ToMetricPrometheusHost(variables),
		PrometheusPort:   env.ToMetricPrometheusPort(variables),
		PrometheusPath:   env.ToMetricPrometheusPath(variables),
	})
	if err != nil {
		tracerCleanup()
		return nil, nil, fmt.Errorf("OpenTelemetry setup failed: %w", err)
	}

	// Set up the MeterProvider to manage metrics
	meterProvider, meterCleanup := intlmetric.NewMeterProvider(ctx, intlmetric.MeterProviderOptions{
		Resource: resource,
		Readers:  metricReaders,
		Logger:   traceLogger,
	})

	// Combine providers into a struct for easy access
	providers := &Providers{
		TracerProvider: tracerProvider,
		MeterProvider:  meterProvider,
	}

	// Return providers, a cleanup function, and no error
	return providers, func() {
		meterCleanup()
		readersCleanup()
		tracerCleanup()
	}, nil
}
