package logs

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/sdk/log"
)

type ExportersOptions struct {
	Exporters   Exporters
	Endpoint    Endpoint
	Compression Compression
	Protocol    Protocol
	Timeout     Timeout
}

func NewExporters(ctx context.Context, options ExportersOptions) ([]log.Exporter, error) {
	exporters := make([]log.Exporter, 0)

	for _, exporter := range options.Exporters {
		switch exporter {
		case ExporterNone:
		case ExporterOtlp:
			exp, err := newExporter(ctx, exporterOptions{
				Endpoint:    options.Endpoint,
				Compression: options.Compression,
				Protocol:    options.Protocol,
				Timeout:     options.Timeout,
			})
			if err != nil {
				return nil, err
			}

			exporters = append(exporters, exp)
		case ExporterConsole:
			exp, err := newConsoleExporter()
			if err != nil {
				return nil, err
			}

			exporters = append(exporters, exp)
		default:
			return nil, fmt.Errorf("unrecognized value for log exporter: %s", exporter)
		}
	}

	return exporters, nil
}

type exporterOptions struct {
	Endpoint    Endpoint
	Compression Compression
	Protocol    Protocol
	Timeout     Timeout
}

func newExporter(ctx context.Context, options exporterOptions) (log.Exporter, error) {
	switch exporter := options.Protocol; exporter {
	case ProtocolGrpc:
		return newGRPCExporter(ctx, grpcExporterOptions{
			Endpoint:    options.Endpoint,
			Compression: options.Compression,
			Timeout:     options.Timeout,
		})
	case ProtocolHttpProtobuf:
		return newHTTPExporter(ctx, httpExporterOptions{
			Endpoint:    options.Endpoint,
			Compression: options.Compression,
			Timeout:     options.Timeout,
		})
	default:
		return nil, fmt.Errorf("unsupported protocol specified for log exporter: %s", exporter)
	}
}

type grpcExporterOptions struct {
	Endpoint    Endpoint
	Compression Compression
	Timeout     Timeout
}

func newGRPCExporter(ctx context.Context, options grpcExporterOptions) (log.Exporter, error) {
	opts := []otlploggrpc.Option{
		otlploggrpc.WithEndpoint(options.Endpoint.Host),
		otlploggrpc.WithTimeout(time.Duration(options.Timeout)),
	}
	if options.Compression == CompressionGzip {
		opts = append(opts, otlploggrpc.WithCompressor(options.Compression.String()))
	}

	if options.Endpoint.Scheme != "https" {
		opts = append(opts, otlploggrpc.WithInsecure())
	}

	exp, err := otlploggrpc.New(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize the OTLP/gRPC Exporter: %w", err)
	}

	return exp, nil
}

type httpExporterOptions struct {
	Endpoint    Endpoint
	Compression Compression
	Timeout     Timeout
}

func newHTTPExporter(ctx context.Context, options httpExporterOptions) (log.Exporter, error) {
	opts := []otlploghttp.Option{
		otlploghttp.WithEndpoint(options.Endpoint.Host),
		otlploghttp.WithTimeout(time.Duration(options.Timeout)),
	}

	switch compression := options.Compression; compression {
	case CompressionNone:
		opts = append(opts, otlploghttp.WithCompression(otlploghttp.NoCompression))
	case CompressionGzip:
		opts = append(opts, otlploghttp.WithCompression(otlploghttp.GzipCompression))
	default:
		return nil, fmt.Errorf("invalid compression for log HTTP exporter: %s", compression)
	}

	exp, err := otlploghttp.New(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("error creating log HTTP exporter: %w", err)
	}

	return exp, nil
}

func newConsoleExporter() (log.Exporter, error) {
	exporter, err := stdoutlog.New(stdoutlog.WithPrettyPrint())
	if err != nil {
		return nil, fmt.Errorf("error creating stdout log exporter: %w", err)
	}

	return exporter, nil
}
