package env

import (
	intlmetric "github.com/angelokurtis/go-otel/starter/internal/metric"
	intltrace "github.com/angelokurtis/go-otel/starter/internal/trace"
)

type ExporterProvider interface {
	TraceExporters() (intltrace.Exporters, bool)
	MetricExporters() (intlmetric.Exporters, bool)
}
