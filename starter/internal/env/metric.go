package env

import (
	"github.com/gotidy/ptr"

	"github.com/angelokurtis/go-otel/starter/internal/metric"
)

func ToMetricExporters(otel *Variables) metric.Exporters {
	return otel.Metrics.Exporter
}

func ToMetricEndpoint(otel *Variables) metric.Endpoint {
	return metric.Endpoint(ptr.ToDef(otel.Exporter.OTLP.Metrics.Endpoint, otel.Exporter.OTLP.Endpoint))
}

func ToMetricCompression(otel *Variables) metric.Compression {
	return ptr.ToDef(otel.Exporter.OTLP.Metrics.Compression, metric.Compression(otel.Exporter.OTLP.Compression))
}

func ToMetricExportInterval(otel *Variables) metric.ExportInterval {
	return metric.ExportInterval(otel.Metric.Export.Interval)
}

func ToMetricProtocol(otel *Variables) metric.Protocol {
	return ptr.ToDef(otel.Exporter.OTLP.Metrics.Protocol, metric.Protocol(otel.Exporter.OTLP.Protocol))
}
