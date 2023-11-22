package otel

import (
	"context"
	"runtime"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// StartSpanFromContext starts a new OpenTelemetry span using the provided context.
// It returns a new context containing the started span and the started span itself.
//
// Example:
//
//	ctx := context.Background()
//	ctx, span := otel.StartSpanFromContext(ctx)
//	defer span.End()
func StartSpanFromContext(ctx context.Context) (context.Context, trace.Span) {
	provider := otel.GetTracerProvider()
	pkg, fn := getCaller()
	tracer := provider.Tracer(pkg)

	return tracer.Start(ctx, fn)
}

func getCaller() (pkg, fn string) {
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
