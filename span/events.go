package span

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

// Event adds an event to the span associated with the provided context.
// The event is described by the given name and may include additional options.
//
// Example:
//
//	ctx, end := span.Start(ctx)
//	defer end()
//	span.Event(ctx, "custom-event", trace.WithAttributes(attribute.Int("count", 42)))
func Event(ctx context.Context, name string, options ...trace.EventOption) {
	span := trace.SpanFromContext(ctx)
	span.AddEvent(name, options...)
}
