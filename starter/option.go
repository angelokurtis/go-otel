package starter

import (
	"github.com/gotidy/ptr"
	"github.com/prometheus/client_golang/prometheus"

	intlmetric "github.com/angelokurtis/go-otel/starter/internal/metric"
	intltrace "github.com/angelokurtis/go-otel/starter/internal/trace"
)

type option struct {
	serviceName        *string
	traceExporters     *intltrace.Exporters
	metricExporters    *intlmetric.Exporters
	prometheusRegistry *prometheus.Registry
}

func newOption(opts ...func(c *option)) *option {
	o := &option{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func (o *option) ServiceName() (string, bool) {
	return ptr.To(o.serviceName), o.serviceName != nil
}

func (o *option) TraceExporters() (intltrace.Exporters, bool) {
	return ptr.To(o.traceExporters), o.traceExporters != nil
}

func (o *option) MetricExporters() (intlmetric.Exporters, bool) {
	return ptr.To(o.metricExporters), o.metricExporters != nil
}

func (o *option) PrometheusRegistry() (*prometheus.Registry, bool) {
	return o.prometheusRegistry, o.prometheusRegistry != nil
}

func WithServiceName(serviceName string) func(opt *option) {
	return func(opt *option) {
		opt.serviceName = ptr.Of(serviceName)
	}
}

func WithTracesExporter(exporter intltrace.Exporter) func(opt *option) {
	return func(opt *option) {
		exporters := ptr.ToDef(opt.traceExporters, make(intltrace.Exporters, 0))
		exporters = append(exporters, exporter)
		opt.traceExporters = ptr.Of(exporters)
	}
}

func WithMetricsExporter(exporter intlmetric.Exporter) func(opt *option) {
	return func(opt *option) {
		exporters := ptr.ToDef(opt.metricExporters, make(intlmetric.Exporters, 0))
		exporters = append(exporters, exporter)
		opt.metricExporters = ptr.Of(exporters)
	}
}

func WithPrometheusRegistry(promRegistry *prometheus.Registry) func(opt *option) {
	return func(opt *option) {
		opt.prometheusRegistry = promRegistry
	}
}
