package config

import (
	"fmt"
	"net/url"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/angelokurtis/go-starter-otel/internal/metric"
	"github.com/angelokurtis/go-starter-otel/internal/trace"
)

// Env provides configuration options for the OpenTelemetry SDK from environment variables.
type Env struct {
	TracesExporter  []trace.Exporter  `envconfig:"TRACES_EXPORTER" default:"otlp"`
	MetricsExporter []metric.Exporter `envconfig:"METRICS_EXPORTER" default:"otlp"`

	ExporterOTLPEndpoint        *url.URL `envconfig:"EXPORTER_OTLP_ENDPOINT" default:"http://localhost:4317"`
	ExporterOTLPTracesEndpoint  *url.URL `envconfig:"EXPORTER_OTLP_TRACES_ENDPOINT"`
	ExporterOTLPMetricsEndpoint *url.URL `envconfig:"EXPORTER_OTLP_METRICS_ENDPOINT"`

	ExporterOTLPTimeout        time.Duration  `envconfig:"EXPORTER_OTLP_TIMEOUT" default:"10s"`
	ExporterOTLPTracesTimeout  *time.Duration `envconfig:"EXPORTER_OTLP_TRACES_TIMEOUT"`
	ExporterOTLPMetricsTimeout *time.Duration `envconfig:"EXPORTER_OTLP_METRICS_TIMEOUT"`

	ExporterOTLPProtocol        Protocol         `envconfig:"EXPORTER_OTLP_PROTOCOL" default:"grpc"`
	ExporterOTLPTracesProtocol  *trace.Protocol  `envconfig:"EXPORTER_OTLP_TRACES_PROTOCOL"`
	ExporterOTLPMetricsProtocol *metric.Protocol `envconfig:"EXPORTER_OTLP_METRICS_PROTOCOL"`

	ExporterOTLPCompression        Compression         `envconfig:"EXPORTER_OTLP_COMPRESSION" default:"gzip"`
	ExporterOTLPTracesCompression  *trace.Compression  `envconfig:"EXPORTER_OTLP_TRACES_COMPRESSION"`
	ExporterOTLPMetricsCompression *metric.Compression `envconfig:"EXPORTER_OTLP_METRICS_COMPRESSION"`

	MetricExportInterval time.Duration `envconfig:"METRIC_EXPORT_INTERVAL" default:"60s"`

	TracesSampler    trace.Sampler `envconfig:"TRACES_SAMPLER" default:"parentbased_always_on"`
	TracesSamplerArg float64       `envconfig:"TRACES_SAMPLER_ARG" default:"0.0"`

	Propagators []trace.Propagator `default:"tracecontext,baggage"`
}

// NewFromEnv takes advantage of the `envconfig` package to parse environment variables into the Env struct. If any of
// the environment variables are missing or have an invalid value, the method returns an error with a detailed message
// that explains the cause of the failure.
func NewFromEnv() (*Env, error) {
	var e Env
	if err := envconfig.Process("OTEL", &e); err != nil {
		return nil, fmt.Errorf("failed to create a new configuration from environment variables: %w", err)
	}

	return &e, nil
}
