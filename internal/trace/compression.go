package trace

import (
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

func newHTTPCompression(config *Config) (otlptracehttp.Compression, error) {
	switch compression := config.Compression; compression {
	case CompressionNone:
		return otlptracehttp.NoCompression, nil
	case CompressionGzip:
		return otlptracehttp.GzipCompression, nil
	default:
		return 0, fmt.Errorf("unrecognized value for traces HTTP compression: %s", compression)
	}
}
