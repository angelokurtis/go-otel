package metric

import (
	"context"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	otelprom "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
)

// ReadersOptions defines configuration for readers, such as exporters and protocols
type ReadersOptions struct {
	Exporters   Exporters
	Endpoint    Endpoint
	Compression Compression
	Protocol    Protocol

	RegistryProvider RegistryProvider

	PrometheusHost PrometheusHost
	PrometheusPort PrometheusPort
	PrometheusPath PrometheusPath
}

// NewReaders creates metric readers for the specified exporters and returns them with shutdown functions
func NewReaders(ctx context.Context, options ReadersOptions) ([]metric.Reader, func(), error) {
	readers := make([]metric.Reader, 0)

	var shutdowns []func()

	// Iterate over each exporter to create corresponding readers
	for _, exporter := range options.Exporters {
		reader, shutdown, err := newReader(ctx, readerOptions{
			Exporter:         exporter,
			Endpoint:         options.Endpoint,
			Compression:      options.Compression,
			Protocol:         options.Protocol,
			RegistryProvider: options.RegistryProvider,
			PrometheusHost:   options.PrometheusHost,
			PrometheusPort:   options.PrometheusPort,
			PrometheusPath:   options.PrometheusPath,
		})
		if err != nil {
			return nil, func() {
				for _, shutdownFn := range shutdowns {
					shutdownFn()
				}
			}, err
		}

		// Add the created reader and shutdown function
		if reader != nil {
			readers = append(readers, reader)
			shutdowns = append(shutdowns, shutdown)
		}
	}

	// Return the readers and a shutdown function
	return readers, func() {
		for _, shutdownFn := range shutdowns {
			shutdownFn()
		}
	}, nil
}

type RegistryProvider interface {
	PrometheusRegistry() (*prometheus.Registry, bool)
}

// readerOptions contains the configuration for a single reader
type readerOptions struct {
	Exporter    Exporter
	Endpoint    Endpoint
	Compression Compression
	Protocol    Protocol

	RegistryProvider RegistryProvider

	PrometheusHost PrometheusHost
	PrometheusPort PrometheusPort
	PrometheusPath PrometheusPath
}

// newReader creates a metric reader for a specific exporter and returns it with a shutdown function
func newReader(ctx context.Context, options readerOptions) (metric.Reader, func(), error) {
	switch exporter := options.Exporter; exporter {
	case ExporterLogging:
		// Create a logging exporter and return the reader with no shutdown function
		exp, err := stdoutmetric.New()
		if err != nil {
			return nil, func() {}, fmt.Errorf("failed to initialize the Logging Exporter: %w", err)
		}

		return metric.NewPeriodicReader(exp), func() {}, nil
	case ExporterOtlp:
		// Create an OTLP exporter and return the reader
		exp, err := newOTLPExporter(ctx, otlpExporterOptions{
			Endpoint:    options.Endpoint,
			Compression: options.Compression,
			Protocol:    options.Protocol,
		})
		if err != nil {
			return nil, func() {}, err
		}

		return metric.NewPeriodicReader(exp), func() {}, nil
	case ExporterPrometheus:
		// Create a Prometheus exporter and return the reader with a shutdown function
		registry := prometheus.NewRegistry()

		reader, err := otelprom.New(
			otelprom.WithRegisterer(registry),
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to initialize the Prometheus Exporter: %w", err)
		}

		exp, err := NewPrometheusServer(
			registry,
			options.PrometheusHost,
			options.PrometheusPort,
			options.PrometheusPath,
		)
		if err != nil {
			return nil, func() {}, err
		}

		shutdown, err := exp.Start(ctx)
		if err != nil {
			return nil, func() {}, err
		}

		return reader, shutdown, nil
	case ExporterNone:
		// No exporter, return nil
		return nil, func() {}, nil
	default:
		// Unknown exporter type
		return nil, func() {}, fmt.Errorf("unrecognized value for metric exporter: %s", exporter)
	}
}
