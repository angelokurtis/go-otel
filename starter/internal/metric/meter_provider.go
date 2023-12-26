package metric

import (
	"context"

	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

type MeterProviderOptions struct {
	Resource *resource.Resource
	Readers  []metric.Reader
	Logger   logr.Logger
}

func NewMeterProvider(ctx context.Context, options MeterProviderOptions) (*metric.MeterProvider, func()) {
	opts := []metric.Option{
		metric.WithResource(options.Resource),
	}

	for _, reader := range options.Readers {
		opts = append(opts, metric.WithReader(reader))
	}

	mp := metric.NewMeterProvider(opts...)
	otel.SetMeterProvider(mp)

	cleanup := func() {
		if err := mp.Shutdown(ctx); err != nil {
			options.Logger.V(1).Info("Error shutting down Meter Provider", "err", err.Error())
		}
	}

	return mp, cleanup
}
