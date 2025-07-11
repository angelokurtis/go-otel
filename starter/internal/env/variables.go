package env

import (
	"fmt"
	"net/url"
	"time"

	env "github.com/caarlos0/env/v11"

	"github.com/angelokurtis/go-otel/starter/internal/logs"
	"github.com/angelokurtis/go-otel/starter/internal/metric"
	"github.com/angelokurtis/go-otel/starter/internal/trace"
)

type Variables struct {
	Traces struct {
		Exporter []trace.Exporter `envDefault:"otlp"`
		Sampler  struct {
			Type trace.Sampler `env:"SAMPLER" envDefault:"parentbased_always_on"`
			Arg  float64       `env:"SAMPLER_ARG" envDefault:"0.0"`
		}
	} `envPrefix:"TRACES_"`
	Metrics struct {
		Exporter []metric.Exporter `envDefault:"otlp"`
	} `envPrefix:"METRICS_"`
	Logs struct {
		Exporter []logs.Exporter `envDefault:"otlp"`
	} `envPrefix:"LOGS_"`
	Exporter struct {
		OTLP struct {
			Endpoint    url.URL       `envDefault:"http://localhost:4317"`
			Timeout     time.Duration `envDefault:"10s"`
			Protocol    Protocol      `envDefault:"grpc"`
			Compression Compression   `envDefault:"gzip"`
			Traces      struct {
				Endpoint    *url.URL
				Timeout     *time.Duration
				Protocol    *trace.Protocol
				Compression *trace.Compression
			} `envPrefix:"TRACES_"`
			Metrics struct {
				Endpoint    *url.URL
				Timeout     *time.Duration
				Protocol    *metric.Protocol
				Compression *metric.Compression
			} `envPrefix:"METRICS_"`
			Logs struct {
				Endpoint    *url.URL
				Timeout     *time.Duration
				Protocol    *logs.Protocol
				Compression *logs.Compression
			} `envPrefix:"LOGS_"`
		} `envPrefix:"EXPORTER_OTLP_"`
		Prometheus struct {
			Host metric.PrometheusHost `envDefault:"0.0.0.0"`
			Port metric.PrometheusPort `envDefault:"9464"`
			Path metric.PrometheusPath `envDefault:"/metrics"`
		}
	}
	Propagators             []trace.Propagator `envDefault:"tracecontext,baggage"`
	BatchLogRecordProcessor struct {
		ScheduleDelay      time.Duration `envDefault:"1s"`
		ExportTimeout      time.Duration `envDefault:"30s"`
		MaxQueueSize       int           `envDefault:"2048"`
		MaxExportBatchSize int           `envDefault:"512"`
	} `envPrefix:"BLRP_"`
}

// LookupVariables takes advantage of the `envconfig` package to parse environment variables into the Variables struct.
// If any of the environment variables are missing or have an invalid value, the method returns an error with a detailed
// message that explains the cause of the failure.
func LookupVariables() (*Variables, error) {
	var otel Variables

	err := env.ParseWithOptions(&otel, env.Options{
		Prefix:                "OTEL_",
		UseFieldNameByDefault: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve OpenTelemetry configuration from environment variables: %w", err)
	}

	return &otel, nil
}
