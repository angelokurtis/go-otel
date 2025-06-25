package logs

import (
	"context"

	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
)

type LoggerProviderOptions struct {
	Resource           *resource.Resource
	Exporters          []log.Exporter
	ExportInterval     ExportInterval
	ExportTimeout      ExportTimeout
	MaxQueueSize       MaxQueueSize
	ExportMaxBatchSize ExportMaxBatchSize
	Logger             logr.Logger
}

func NewLoggerProvider(ctx context.Context, options LoggerProviderOptions) (*log.LoggerProvider, func(), error) {
	opts := []log.LoggerProviderOption{
		log.WithResource(options.Resource),
	}

	for _, exporter := range options.Exporters {
		var processor log.Processor

		switch exp := exporter.(type) {
		case *stdoutlog.Exporter:
			processor = log.NewSimpleProcessor(exp)
		default:
			processor = newBatchProcessor(ctx, batchProcessorOptions{
				Exporter:           exporter,
				ExportInterval:     options.ExportInterval,
				ExportTimeout:      options.ExportTimeout,
				MaxQueueSize:       options.MaxQueueSize,
				ExportMaxBatchSize: options.ExportMaxBatchSize,
			})
		}

		opts = append(opts, log.WithProcessor(processor))
	}

	prov := log.NewLoggerProvider(opts...)

	cleanup := func() {
		if err := prov.Shutdown(ctx); err != nil {
			options.Logger.V(1).Info("Error shutting down Logger Provider", "err", err.Error())
		}
	}

	return prov, cleanup, nil
}
