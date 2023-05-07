package trace

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func NewTracerProvider(res *resource.Resource, sampler trace.Sampler, exporters []trace.SpanExporter, config Config) *trace.TracerProvider {
	options := []trace.TracerProviderOption{
		trace.WithResource(res),
		trace.WithSampler(sampler),
	}
	for _, exp := range exporters {
		options = append(options, trace.WithBatcher(
			exp,
			trace.WithBatchTimeout(config.ExporterOTLPTracesTimeout())),
		)
	}

	tp := trace.NewTracerProvider(options...)
	otel.SetTracerProvider(tp)

	return tp
}

func NewSpanExporters(ctx context.Context, config Config, client otlptrace.Client) ([]trace.SpanExporter, error) {
	exporters := make([]trace.SpanExporter, 0)

	for _, expType := range config.TracesExporter() {
		exp, err := newSpanExporter(ctx, expType, client)
		if err != nil {
			return nil, err
		}

		exporters = append(exporters, exp)
	}

	return exporters, nil
}

func NewSampler(ctx context.Context, config Config) (trace.Sampler, error) {
	switch sampler := config.TracesSampler(); sampler {
	case SamplerAlwaysOn:
		return trace.AlwaysSample(), nil
	case SamplerAlwaysOff:
		return trace.NeverSample(), nil
	case SamplerTraceidratio:
		fraction := config.TracesSamplerArg()
		return trace.TraceIDRatioBased(fraction), nil
	case SamplerParentbasedAlwaysOn:
		return trace.ParentBased(trace.AlwaysSample()), nil
	case SamplerParentbasedAlwaysOff:
		return trace.ParentBased(trace.NeverSample()), nil
	case SamplerParentbasedTraceidratio:
		fraction := config.TracesSamplerArg()
		return trace.ParentBased(trace.TraceIDRatioBased(fraction)), nil
	default:
		return nil, fmt.Errorf("unsupported sampler option: %s", sampler)
	}
}

func NewTraceClient(config Config) (otlptrace.Client, error) {
	switch protocol := config.ExporterOTLPTracesProtocol(); protocol {
	case ProtocolGrpc:
		return newGRPCTraceClient(config)
	case ProtocolHttpProtobuf:
		return newHTTPTraceClient(config)
	default:
		return nil, fmt.Errorf("unsupported OTLP traces protocol: %s", protocol)
	}
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

func newGRPCTraceClient(config Config) (otlptrace.Client, error) {
	endpoint := config.ExporterOTLPTracesEndpoint()
	opts := []otlptracegrpc.Option{otlptracegrpc.WithEndpoint(endpoint.Host)}

	if endpoint.Scheme != "https" {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}

	return otlptracegrpc.NewClient(opts...), nil
}

func newHTTPTraceClient(config Config) (otlptrace.Client, error) {
	compression, err := newHTTPCompression(config)
	if err != nil {
		return nil, err
	}

	endpoint := config.ExporterOTLPTracesEndpoint()
	opts := []otlptracehttp.Option{
		otlptracehttp.WithCompression(compression),
		otlptracehttp.WithEndpoint(endpoint.Host),
	}

	if endpoint.Scheme != "https" {
		opts = append(opts, otlptracehttp.WithInsecure())
	}

	return otlptracehttp.NewClient(opts...), nil
}

func newHTTPCompression(config Config) (otlptracehttp.Compression, error) {
	switch compression := config.ExporterOTLPTracesCompression(); compression {
	case CompressionNone:
		return otlptracehttp.NoCompression, nil
	case CompressionGzip:
		return otlptracehttp.GzipCompression, nil
	default:
		return 0, fmt.Errorf("unrecognized value for traces HTTP compression: %s", compression)
	}
}
