package span

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Attributes sets attributes on the span extracted from the given context.
// If the context does not contain a span, this function has no effect.
//
// Example:
//
//	span.Attributes(ctx, attribute.String("key", "value"), attribute.Int("count", 42))
func Attributes(ctx context.Context, attributes ...attribute.KeyValue) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(attributes...)
}
