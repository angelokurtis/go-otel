package env

import (
	"github.com/gotidy/ptr"

	"github.com/angelokurtis/go-otel/starter/internal/metric"
)

func ToMetricExporters(otel *Variables, provider ExporterProvider) metric.Exporters {
	if exporters, ok := provider.MetricExporters(); ok {
		return exporters
	}

	return otel.Metrics.Exporter
}

func ToMetricEndpoint(otel *Variables) metric.Endpoint {
	return metric.Endpoint(ptr.ToDef(otel.Exporter.OTLP.Metrics.Endpoint, otel.Exporter.OTLP.Endpoint))
}

func ToMetricCompression(otel *Variables) metric.Compression {
	return ptr.ToDef(otel.Exporter.OTLP.Metrics.Compression, metric.Compression(otel.Exporter.OTLP.Compression))
}

func ToMetricProtocol(otel *Variables) metric.Protocol {
	return ptr.ToDef(otel.Exporter.OTLP.Metrics.Protocol, metric.Protocol(otel.Exporter.OTLP.Protocol))
}

func ToMetricPrometheusHost(otel *Variables) metric.PrometheusHost {
	return otel.Exporter.Prometheus.Host
}

func ToMetricPrometheusPort(otel *Variables) metric.PrometheusPort {
	return otel.Exporter.Prometheus.Port
}

func ToMetricPrometheusPath(otel *Variables) metric.PrometheusPath {
	return otel.Exporter.Prometheus.Path
}
