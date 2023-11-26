package trace

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

type SpanExportersOptions struct {
	Exporters Exporters
	Client    otlptrace.Client
}

func NewSpanExporters(ctx context.Context, options SpanExportersOptions) ([]trace.SpanExporter, error) {
	exporters := make([]trace.SpanExporter, 0)

	for _, exporter := range options.Exporters {
		exp, err := newSpanExporter(ctx, spanExporterOptions{
			Exporter: exporter,
			Client:   options.Client,
		})
		if err != nil {
			return nil, err
		}

		exporters = append(exporters, exp)
	}

	return exporters, nil
}

type spanExporterOptions struct {
	Exporter Exporter
	Client   otlptrace.Client
}

func newSpanExporter(ctx context.Context, options spanExporterOptions) (trace.SpanExporter, error) {
	switch exporter := options.Exporter; exporter {
	case ExporterNone:
		return tracetest.NewNoopExporter(), nil
	case ExporterOtlp:
		exp, err := otlptrace.New(ctx, options.Client)
		if err != nil {
			return nil, fmt.Errorf("an error occurred while creating a new span exporter: %w", err)
		}

		return exp, nil
	default:
		return nil, fmt.Errorf("unrecognized value for traces exporter: %s", exporter)
	}
}
