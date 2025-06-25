package logs

import (
	"time"

	"go.opentelemetry.io/otel/sdk/log"
	"golang.org/x/net/context"
)

type batchProcessorOptions struct {
	Exporter           log.Exporter
	ExportInterval     ExportInterval
	ExportTimeout      ExportTimeout
	MaxQueueSize       MaxQueueSize
	ExportMaxBatchSize ExportMaxBatchSize
}

func newBatchProcessor(ctx context.Context, options batchProcessorOptions) *log.BatchProcessor {
	processor := log.NewBatchProcessor(options.Exporter,
		log.WithExportInterval(time.Duration(options.ExportInterval)),
		log.WithExportTimeout(time.Duration(options.ExportTimeout)),
		log.WithMaxQueueSize(int(options.MaxQueueSize)),
		log.WithExportMaxBatchSize(int(options.ExportMaxBatchSize)),
	)

	return processor
}
