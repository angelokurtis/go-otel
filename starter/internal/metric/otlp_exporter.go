package metric

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/sdk/metric"
)

type otlpExporterOptions struct {
	Endpoint    Endpoint
	Compression Compression
	Protocol    Protocol
}

func newOTLPExporter(ctx context.Context, options otlpExporterOptions) (metric.Exporter, error) {
	switch exporter := options.Protocol; exporter {
	case ProtocolGrpc:
		return newGRPCExporter(ctx, grpcExporterOptions{
			Endpoint:    options.Endpoint,
			Compression: options.Compression,
		})
	case ProtocolHttpProtobuf:
		return newHTTPExporter(ctx, httpExporterOptions{
			Endpoint:    options.Endpoint,
			Compression: options.Compression,
		})
	default:
		return nil, fmt.Errorf("unrecognized value for metric exporter protocol: %s", exporter)
	}
}

type grpcExporterOptions struct {
	Endpoint    Endpoint
	Compression Compression
}

func newGRPCExporter(ctx context.Context, options grpcExporterOptions) (metric.Exporter, error) {
	endpoint := options.Endpoint

	opts := []otlpmetricgrpc.Option{
		otlpmetricgrpc.WithEndpoint(endpoint.Host),
	}
	if options.Compression == CompressionGzip {
		opts = append(opts, otlpmetricgrpc.WithCompressor(options.Compression.String()))
	}

	if endpoint.Scheme != "https" {
		opts = append(opts, otlpmetricgrpc.WithInsecure())
	}

	// TODO: add config for "OTEL_EXPORTER_OTLP_METRICS_TEMPORALITY_PREFERENCE"
	// TODO: add config for "OTEL_EXPORTER_OTLP_METRICS_DEFAULT_HISTOGRAM_AGGREGATION"

	exp, err := otlpmetricgrpc.New(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize the OTLP/gRPC Exporter: %w", err)
	}

	return exp, nil
}

type httpExporterOptions struct {
	Endpoint    Endpoint
	Compression Compression
}

func newHTTPExporter(ctx context.Context, options httpExporterOptions) (metric.Exporter, error) {
	endpoint := options.Endpoint

	opts := []otlpmetrichttp.Option{
		otlpmetrichttp.WithEndpoint(endpoint.Host),
	}

	switch compression := options.Compression; compression {
	case CompressionNone:
		opts = append(opts, otlpmetrichttp.WithCompression(otlpmetrichttp.NoCompression))
	case CompressionGzip:
		opts = append(opts, otlpmetrichttp.WithCompression(otlpmetrichttp.GzipCompression))
	default:
		return nil, fmt.Errorf("unrecognized value for traces GRPC compression: %s", compression)
	}

	if endpoint.Scheme != "https" {
		opts = append(opts, otlpmetrichttp.WithInsecure())
	}

	// TODO: add config for "OTEL_EXPORTER_OTLP_METRICS_TEMPORALITY_PREFERENCE"
	// TODO: add config for "OTEL_EXPORTER_OTLP_METRICS_DEFAULT_HISTOGRAM_AGGREGATION"

	exp, err := otlpmetrichttp.New(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize the OTLP/HTTP Exporter: %w", err)
	}

	return exp, nil
}
