package env

import (
	intllogs "github.com/angelokurtis/go-otel/starter/internal/logs"
	intlmetric "github.com/angelokurtis/go-otel/starter/internal/metric"
	intltrace "github.com/angelokurtis/go-otel/starter/internal/trace"
)

type ExporterProvider interface {
	TraceExporters() (intltrace.Exporters, bool)
	MetricExporters() (intlmetric.Exporters, bool)
	LogExporters() (intllogs.Exporters, bool)
}
