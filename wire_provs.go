package otel

import (
	"github.com/google/wire"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/angelokurtis/go-starter-otel/internal/env"
	"github.com/angelokurtis/go-starter-otel/internal/logger"
	intlmetric "github.com/angelokurtis/go-starter-otel/internal/metric"
	intltrace "github.com/angelokurtis/go-starter-otel/internal/trace"
)

//nolint:unused // This function is used during compile-time to generate code for dependency injection
var providers = wire.NewSet(
	env.LookupVariables,
	env.ToMetricCompression,
	env.ToMetricEndpoint,
	env.ToMetricExporters,
	env.ToMetricExportInterval,
	env.ToMetricProtocol,
	env.ToTraceCompression,
	env.ToTraceEndpoint,
	env.ToTraceExporters,
	env.ToTracePropagators,
	env.ToTraceProtocol,
	env.ToTraceSampler,
	env.ToTraceSamplerArg,
	env.ToTraceTimeout,
	intlmetric.NewMeterProvider,
	intlmetric.NewReaders,
	intltrace.NewClient,
	intltrace.NewResource,
	intltrace.NewSampler,
	intltrace.NewSpanExporters,
	intltrace.NewTextMapPropagator,
	intltrace.NewTracerProvider,
	logger.New,
	wire.Struct(new(intlmetric.MeterProviderOptions), "*"),
	wire.Struct(new(intlmetric.ReadersOptions), "*"),
	wire.Struct(new(intltrace.ClientOptions), "*"),
	wire.Struct(new(intltrace.SamplerOptions), "*"),
	wire.Struct(new(intltrace.SpanExportersOptions), "*"),
	wire.Struct(new(intltrace.TracerProviderOptions), "*"),
	wire.Struct(new(Manager), "*"),
)

type Manager struct {
	TracerProvider *sdktrace.TracerProvider
	MeterProvider  *sdkmetric.MeterProvider
}
