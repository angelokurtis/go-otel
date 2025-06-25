package starter

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/angelokurtis/go-otel/starter/internal/env"
	"github.com/angelokurtis/go-otel/starter/internal/logger"
	intllog "github.com/angelokurtis/go-otel/starter/internal/logs"
	intlmetric "github.com/angelokurtis/go-otel/starter/internal/metric"
	intltrace "github.com/angelokurtis/go-otel/starter/internal/trace"
)

// Providers holds OpenTelemetry providers for tracing, metrics, and logs.
type Providers struct {
	TracerProvider *trace.TracerProvider
	MeterProvider  *metric.MeterProvider
	LoggerProvider *log.LoggerProvider
}

// ShutdownFunc represents a function used to release all allocated resources.
type ShutdownFunc func()

// StartProviders initializes OpenTelemetry providers for tracing, metrics, and logging.
// It returns the initialized Providers, a shutdown function, and an error (if any occurred).
func StartProviders(ctx context.Context, opts ...func(c *option)) (*Providers, ShutdownFunc, error) {
	// Load user-defined or default options
	options := newOption(opts...)

	// Create a resource to identify the telemetry source
	resource, err := intltrace.NewResource(ctx, options)
	if err != nil {
		return nil, nil, fmt.Errorf("OpenTelemetry setup failed: %w", err)
	}

	// Load configuration from environment variables
	vars, err := env.LookupVariables()
	if err != nil {
		return nil, nil, fmt.Errorf("OpenTelemetry setup failed: %w", err)
	}

	// --- TRACING SETUP ---

	// Configure trace sampler
	sampler, err := intltrace.NewSampler(intltrace.SamplerOptions{
		Sampler:    env.ToTraceSampler(vars),
		SamplerArg: env.ToTraceSamplerArg(vars),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("OpenTelemetry setup failed: %w", err)
	}

	// Create a client for trace exporters
	traceClient, err := intltrace.NewClient(intltrace.ClientOptions{
		Protocol:    env.ToTraceProtocol(vars),
		Endpoint:    env.ToTraceEndpoint(vars),
		Compression: env.ToTraceCompression(vars),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("OpenTelemetry setup failed: %w", err)
	}

	// Configure span exporters
	spanExporters, err := intltrace.NewSpanExporters(ctx, intltrace.SpanExportersOptions{
		Exporters: env.ToTraceExporters(vars, options),
		Client:    traceClient,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("OpenTelemetry setup failed: %w", err)
	}

	// Create a logger for trace-related diagnostics
	traceLogger := logger.New()

	// Initialize the tracer provider
	propagator := intltrace.NewTextMapPropagator(env.ToTracePropagators(vars))
	tracerProvider, tracerCleanup := intltrace.NewTracerProvider(ctx, intltrace.TracerProviderOptions{
		Resource:     resource,
		Sampler:      sampler,
		Exporters:    spanExporters,
		BatchTimeout: env.ToTraceTimeout(vars),
		Propagator:   propagator,
		Logger:       traceLogger,
	})

	// --- METRICS SETUP ---

	// Create metric readers (including Prometheus if configured)
	metricReaders, metricReadersCleanup, err := intlmetric.NewReaders(ctx, intlmetric.ReadersOptions{
		Exporters:          env.ToMetricExporters(vars, options),
		Endpoint:           env.ToMetricEndpoint(vars),
		Compression:        env.ToMetricCompression(vars),
		Protocol:           env.ToMetricProtocol(vars),
		PrometheusProvider: options,
		PrometheusHost:     env.ToMetricPrometheusHost(vars),
		PrometheusPort:     env.ToMetricPrometheusPort(vars),
		PrometheusPath:     env.ToMetricPrometheusPath(vars),
	})
	if err != nil {
		tracerCleanup()
		return nil, nil, fmt.Errorf("OpenTelemetry setup failed: %w", err)
	}

	// Initialize the meter provider for metrics
	meterProvider, meterCleanup := intlmetric.NewMeterProvider(ctx, intlmetric.MeterProviderOptions{
		Resource: resource,
		Readers:  metricReaders,
		Logger:   traceLogger,
	})

	// --- LOGGING SETUP ---

	// Set up exporters for structured logs
	logExporters := env.ToLogExporters(vars, options)

	logExporterOpts := intllog.ExportersOptions{
		Exporters:   logExporters,
		Endpoint:    env.ToLogEndpoint(vars),
		Compression: env.ToLogCompression(vars),
		Protocol:    env.ToLogProtocol(vars),
		Timeout:     env.ToLogTimeout(vars),
	}

	logExporter, err := intllog.NewExporters(ctx, logExporterOpts)
	if err != nil {
		meterCleanup()
		metricReadersCleanup()
		tracerCleanup()

		return nil, nil, fmt.Errorf("OpenTelemetry setup failed: %w", err)
	}

	exportMaxBatchSize, err := env.ToLogMaxExportBatchSize(vars)
	if err != nil {
		meterCleanup()
		metricReadersCleanup()
		tracerCleanup()

		return nil, nil, fmt.Errorf("OpenTelemetry setup failed: %w", err)
	}

	// Initialize the logger provider
	loggerProvider, loggerCleanup, err := intllog.NewLoggerProvider(ctx, intllog.LoggerProviderOptions{
		Resource:           resource,
		Exporters:          logExporter,
		ExportInterval:     env.ToLogScheduleDelay(vars),
		ExportTimeout:      env.ToLogExportTimeout(vars),
		MaxQueueSize:       env.ToLogMaxQueueSize(vars),
		ExportMaxBatchSize: exportMaxBatchSize,
		Logger:             traceLogger,
	})
	if err != nil {
		meterCleanup()
		metricReadersCleanup()
		tracerCleanup()

		return nil, nil, fmt.Errorf("OpenTelemetry setup failed: %w", err)
	}

	// Return the initialized providers and a single unified cleanup function
	providers := &Providers{
		TracerProvider: tracerProvider,
		MeterProvider:  meterProvider,
		LoggerProvider: loggerProvider,
	}

	return providers, func() {
		loggerCleanup()
		meterCleanup()
		metricReadersCleanup()
		tracerCleanup()
	}, nil
}
