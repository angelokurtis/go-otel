package span

import (
	"context"
	"runtime"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const (
	libraryName    = "github.com/angelokurtis/go-otel/span"
	libraryVersion = "v1.2.5"
)

// EndFunc is a function type that can be used to end an OpenTelemetry span.
type EndFunc func(options ...trace.SpanEndOption)

// Start starts a new OpenTelemetry span using the provided context.
// It returns a new context containing the started span and a function to end the span.
//
// Example:
//
//	ctx, end := span.Start(ctx)
//	defer end()
func Start(ctx context.Context, options ...trace.SpanStartOption) (context.Context, EndFunc) {
	provider := otel.GetTracerProvider()
	_, fn := getCaller()
	tracer := provider.Tracer(libraryName, trace.WithInstrumentationVersion(libraryVersion))

	ctx, span := tracer.Start(ctx, fn, options...)

	return ctx, span.End
}

// StartWithName starts a new OpenTelemetry span with the given name using the provided context.
// It returns a new context containing the started span and a function to end the span.
func StartWithName(ctx context.Context, name string, options ...trace.SpanStartOption) (context.Context, EndFunc) {
	provider := otel.GetTracerProvider()
	tracer := provider.Tracer(libraryName, trace.WithInstrumentationVersion(libraryVersion))

	ctx, span := tracer.Start(ctx, name, options...)

	return ctx, span.End
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
