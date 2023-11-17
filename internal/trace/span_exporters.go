package trace

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func NewSpanExporters(ctx context.Context, config *Config, client otlptrace.Client) ([]trace.SpanExporter, error) {
	exporters := make([]trace.SpanExporter, 0)

	for _, expType := range config.Exporters {
		exp, err := newSpanExporter(ctx, expType, client)
		if err != nil {
			return nil, err
		}

		exporters = append(exporters, exp)
	}

	return exporters, nil
}

func newSpanExporter(ctx context.Context, exporter Exporter, client otlptrace.Client) (trace.SpanExporter, error) {
	switch exporter {
	case ExporterNone:
		return tracetest.NewNoopExporter(), nil
	case ExporterOtlp:
		exp, err := otlptrace.New(ctx, client)
		if err != nil {
			return nil, fmt.Errorf("an error occurred while creating a new span exporter: %w", err)
		}

		return exp, nil
	default:
		return nil, fmt.Errorf("unrecognized value for traces exporter: %s", exporter)
	}
}
