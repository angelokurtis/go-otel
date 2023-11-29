package metric

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
)

type ReadersOptions struct {
	Exporters      Exporters
	Endpoint       Endpoint
	Compression    Compression
	ExportInterval ExportInterval
	Protocol       Protocol
}

func NewReaders(ctx context.Context, options ReadersOptions) ([]metric.Reader, error) {
	readers := make([]metric.Reader, 0)

	for _, exporter := range options.Exporters {
		reader, err := newReader(ctx, readerOptions{
			Exporter:       exporter,
			Endpoint:       options.Endpoint,
			Compression:    options.Compression,
			ExportInterval: options.ExportInterval,
			Protocol:       options.Protocol,
		})
		if err != nil {
			return nil, err
		}

		if reader == nil {
			continue
		}

		readers = append(readers, reader)
	}

	return readers, nil
}

type readerOptions struct {
	Exporter       Exporter
	Endpoint       Endpoint
	Compression    Compression
	ExportInterval ExportInterval
	Protocol       Protocol
}

func newReader(ctx context.Context, options readerOptions) (metric.Reader, error) {
	switch exporter := options.Exporter; exporter {
	case ExporterLogging:
		exp, err := stdoutmetric.New()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize the Logging Exporter: %w", err)
		}

		return metric.NewPeriodicReader(
			exp,
			metric.WithInterval(time.Duration(options.ExportInterval)),
		), nil
	case ExporterOtlp:
		exp, err := newOTLPExporter(ctx, otlpExporterOptions{
			Endpoint:       options.Endpoint,
			Compression:    options.Compression,
			ExportInterval: options.ExportInterval,
			Protocol:       options.Protocol,
		})
		if err != nil {
			return nil, err
		}

		return metric.NewPeriodicReader(
			exp,
			metric.WithInterval(time.Duration(options.ExportInterval)),
		), nil
	case ExporterPrometheus:
		reader, err := prometheus.New()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize the Prometheus Exporter: %w", err)
		}
		// TODO: start listening on :9090/metrics
		return reader, nil
	case ExporterNone:
		return nil, nil
	default:
		return nil, fmt.Errorf("unrecognized value for metric exporter: %s", exporter)
	}
}
