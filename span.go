package otel

import (
	"context"
	"runtime"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

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
