package otel

import (
	"context"
	"runtime"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	"github.com/angelokurtis/go-starter-otel/internal/env"
	intltrace "github.com/angelokurtis/go-starter-otel/internal/trace"
)

type Providers struct {
	TracerProvider *sdktrace.TracerProvider

	propagation propagation.TextMapPropagator
}

func SetupProviders(ctx context.Context) (*Providers, func(), error) {
	r := resource.Environment()

	o, err := env.LookupOTel()
	if err != nil {
		return nil, nil, err
	}

	config := env.ToTrace(o)

	s, err := intltrace.NewSampler(config)
	if err != nil {
		return nil, nil, err
	}

	c, err := intltrace.NewTraceClient(config)
	if err != nil {
		return nil, nil, err
	}

	se, err := intltrace.NewSpanExporters(ctx, config, c)
	if err != nil {
		return nil, nil, err
	}

	tp, cleanup := intltrace.NewTracerProvider(ctx, r, s, se, config)
	p := intltrace.NewTextMapPropagator(config)
	provider := &Providers{
		TracerProvider: tp,
		propagation:    p,
	}

	return provider, func() {
		cleanup()
	}, nil
}

func StartSpanFromContext(ctx context.Context) (context.Context, trace.Span) {
	provider := otel.GetTracerProvider()
	pkg, fn := getCaller()
	tracer := provider.Tracer(pkg)

	return tracer.Start(ctx, fn)
}

func getCaller() (string, string) {
	pc, _, _, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc).Name()

	return getPackage(f), getFunction(f)
}

func getPackage(input string) string {
	i1 := strings.LastIndex(input, "/")
	remaining := input[i1+1:]

	i2 := strings.Index(remaining, ".")
	if i2 == -1 {
		return input
	}

	return input[:i1+1+i2]
}

func getFunction(input string) string {
	index := strings.LastIndex(input, "/")
	if index == -1 {
		return input
	}

	return input[index+1:]
}
