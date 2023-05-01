package config

import (
	"fmt"
	"net/url"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Env provides configuration options for the OpenTelemetry SDK from environment variables.
type Env struct {
	TracesExporter  string `default:"otlp"`
	MetricsExporter string `default:"otlp"`
	LogsExporter    string `default:"none"`

	ExporterOTLPTracesEndpoint  *url.URL `default:"http://localhost:4317"`
	ExporterOTLPMetricsEndpoint *url.URL `default:"http://localhost:4317"`
	ExporterOTLPLogsEndpoint    *url.URL `default:"http://localhost:4317"`

	ExporterOTLPTracesTimeout  time.Duration `default:"10s"`
	ExporterOTLPMetricsTimeout time.Duration `default:"10s"`
	ExporterOTLPLogsTimeout    time.Duration `default:"10s"`

	ExporterOTLPTracesProtocol  string `default:"grpc"`
	ExporterOTLPMetricsProtocol string `default:"grpc"`
	ExporterOTLPLogsProtocol    string `default:"grpc"`

	MetricExportInterval time.Duration `default:"60s"`

	Propagators []string `default:"tracecontext,baggage"`
}

// NewFromEnv takes advantage of the `envconfig` package to parse environment variables into the Env struct. If any of
// the environment variables are missing or have an invalid value, the method returns an error with a detailed message
// that explains the cause of the failure.
func NewFromEnv() (*Env, error) {
	var e Env
	if err := envconfig.Process("otel", &e); err != nil {
		return nil, fmt.Errorf("failed to create a new configuration from environment variables: %w", err)
	}

	return &e, nil
}
