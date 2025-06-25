package starter

import (
	"github.com/gotidy/ptr"
	"github.com/prometheus/client_golang/prometheus"

	intllogs "github.com/angelokurtis/go-otel/starter/internal/logs"
	intlmetric "github.com/angelokurtis/go-otel/starter/internal/metric"
	intltrace "github.com/angelokurtis/go-otel/starter/internal/trace"
)

type option struct {
	serviceName          *string
	traceExporters       *intltrace.Exporters
	metricExporters      *intlmetric.Exporters
	logExporters         *intllogs.Exporters
	prometheusGatherer   prometheus.Gatherer
	prometheusRegisterer prometheus.Registerer
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

func (o *option) LogExporters() (intllogs.Exporters, bool) {
	return ptr.To(o.logExporters), o.logExporters != nil
}

func (o *option) PrometheusGatherer() (prometheus.Gatherer, bool) {
	return o.prometheusGatherer, o.prometheusGatherer != nil
}

func (o *option) PrometheusRegisterer() (prometheus.Registerer, bool) {
	return o.prometheusRegisterer, o.prometheusRegisterer != nil
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

func WithLogsExporter(exporter intllogs.Exporter) func(opt *option) {
	return func(opt *option) {
		exporters := ptr.ToDef(opt.logExporters, make(intllogs.Exporters, 0))
		exporters = append(exporters, exporter)
		opt.logExporters = ptr.Of(exporters)
	}
}

func WithPrometheusGatherer(gatherer prometheus.Gatherer) func(opt *option) {
	return func(opt *option) {
		opt.prometheusGatherer = gatherer
	}
}

func WithPrometheusRegisterer(registerer prometheus.Registerer) func(opt *option) {
	return func(opt *option) {
		opt.prometheusRegisterer = registerer
	}
}
