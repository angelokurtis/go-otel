package trace

import (
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

func NewTraceClient(config *Config) (otlptrace.Client, error) {
	switch protocol := config.Protocol; protocol {
	case ProtocolGrpc:
		return newGRPCTraceClient(config)
	case ProtocolHttpProtobuf:
		return newHTTPTraceClient(config)
	default:
		return nil, fmt.Errorf("unsupported OTLP traces protocol: %s", protocol)
	}
}

func newGRPCTraceClient(config *Config) (otlptrace.Client, error) {
	endpoint := config.Endpoint
	opts := []otlptracegrpc.Option{otlptracegrpc.WithEndpoint(endpoint.Host)}

	if endpoint.Scheme != "https" {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}

	return otlptracegrpc.NewClient(opts...), nil
}

func newHTTPTraceClient(config *Config) (otlptrace.Client, error) {
	compression, err := newHTTPCompression(config)
	if err != nil {
		return nil, err
	}

	endpoint := config.Endpoint
	opts := []otlptracehttp.Option{
		otlptracehttp.WithCompression(compression),
		otlptracehttp.WithEndpoint(endpoint.Host),
	}

	if endpoint.Scheme != "https" {
		opts = append(opts, otlptracehttp.WithInsecure())
	}

	return otlptracehttp.NewClient(opts...), nil
}
