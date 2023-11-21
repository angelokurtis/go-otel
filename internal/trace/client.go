package trace

import (
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

type ClientOptions struct {
	Protocol    Protocol
	Endpoint    Endpoint
	Compression Compression
}

func NewClient(options ClientOptions) (otlptrace.Client, error) {
	switch protocol := options.Protocol; protocol {
	case ProtocolGrpc:
		return newGRPCClient(grpcClientOptions{
			Endpoint:    options.Endpoint,
			Compression: options.Compression,
		})
	case ProtocolHttpProtobuf:
		return newHTTPClient(httpClientOptions{
			Endpoint:    options.Endpoint,
			Compression: options.Compression,
		})
	default:
		return nil, fmt.Errorf("unsupported OTLP traces protocol: %s", protocol)
	}
}

type grpcClientOptions struct {
	Endpoint    Endpoint
	Compression Compression
}

func newGRPCClient(options grpcClientOptions) (otlptrace.Client, error) {
	endpoint := options.Endpoint
	opts := []otlptracegrpc.Option{otlptracegrpc.WithEndpoint(endpoint.Host)}

	switch compression := options.Compression; compression {
	case CompressionNone:
	case CompressionGzip:
		opts = append(opts, otlptracegrpc.WithCompressor(compression.String()))
	default:
		return nil, fmt.Errorf("unrecognized value for traces GRPC compression: %s", compression)
	}

	if endpoint.Scheme != "https" {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}

	return otlptracegrpc.NewClient(opts...), nil
}

type httpClientOptions struct {
	Endpoint    Endpoint
	Compression Compression
}

func newHTTPClient(options httpClientOptions) (otlptrace.Client, error) {
	endpoint := options.Endpoint
	opts := []otlptracehttp.Option{otlptracehttp.WithEndpoint(endpoint.Host)}

	switch compression := options.Compression; compression {
	case CompressionNone:
		opts = append(opts, otlptracehttp.WithCompression(otlptracehttp.NoCompression))
	case CompressionGzip:
		opts = append(opts, otlptracehttp.WithCompression(otlptracehttp.GzipCompression))
	default:
		return nil, fmt.Errorf("unrecognized value for traces HTTP compression: %s", compression)
	}

	if endpoint.Scheme != "https" {
		opts = append(opts, otlptracehttp.WithInsecure())
	}

	return otlptracehttp.NewClient(opts...), nil
}
