package env

import (
	"fmt"

	"github.com/gotidy/ptr"

	"github.com/angelokurtis/go-otel/starter/internal/logs"
)

func ToLogExporters(otel *Variables, provider ExporterProvider) logs.Exporters {
	if exp, ok := provider.LogExporters(); ok {
		return exp
	}

	return otel.Logs.Exporter
}

func ToLogEndpoint(otel *Variables) logs.Endpoint {
	ep := otel.Exporter.OTLP.Logs.Endpoint
	def := otel.Exporter.OTLP.Endpoint

	return logs.Endpoint(ptr.ToDef(ep, def))
}

func ToLogTimeout(otel *Variables) logs.Timeout {
	t := otel.Exporter.OTLP.Logs.Timeout
	def := otel.Exporter.OTLP.Timeout

	return logs.Timeout(ptr.ToDef(t, def))
}

func ToLogProtocol(otel *Variables) logs.Protocol {
	p := otel.Exporter.OTLP.Logs.Protocol
	def := logs.Protocol(otel.Exporter.OTLP.Protocol)

	return ptr.ToDef(p, def)
}

func ToLogCompression(otel *Variables) logs.Compression {
	c := otel.Exporter.OTLP.Logs.Compression
	def := logs.Compression(otel.Exporter.OTLP.Compression)

	return ptr.ToDef(c, def)
}

func ToLogScheduleDelay(otel *Variables) logs.ExportInterval {
	return logs.ExportInterval(otel.BatchLogRecordProcessor.ScheduleDelay)
}

func ToLogExportTimeout(otel *Variables) logs.ExportTimeout {
	return logs.ExportTimeout(otel.BatchLogRecordProcessor.ExportTimeout)
}

func ToLogMaxQueueSize(otel *Variables) logs.MaxQueueSize {
	return logs.MaxQueueSize(otel.BatchLogRecordProcessor.MaxQueueSize)
}

func ToLogMaxExportBatchSize(otel *Variables) (logs.ExportMaxBatchSize, error) {
	batch := otel.BatchLogRecordProcessor.MaxExportBatchSize
	queue := otel.BatchLogRecordProcessor.MaxQueueSize

	if batch <= 0 {
		return 0, fmt.Errorf("OTEL_BLRP_MAX_EXPORT_BATCH_SIZE must be positive: got %d", batch)
	}

	if batch > queue {
		return 0, fmt.Errorf("OTEL_BLRP_MAX_EXPORT_BATCH_SIZE (%d) must be ≤ OTEL_BLRP_MAX_QUEUE_SIZE (%d)", batch, queue)
	}

	return logs.ExportMaxBatchSize(batch), nil
}
