package span

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// WrapError wraps the given error and records it as an error in the current OpenTelemetry span.
// If the provided error is nil, it returns nil.
func WrapError(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}

	span := trace.SpanFromContext(ctx)

	uerr := errors.Unwrap(err)
	if uerr == nil {
		uerr = err
	}

	// Record the error and set its status in the span.
	span.RecordError(uerr)
	span.SetStatus(codes.Error, uerr.Error())

	return err
}

// WrapErrorWithStack wraps the given error, records it as an error in the current OpenTelemetry span,
// and includes the stack trace information in the recorded error.
// If the provided error is nil, it returns nil.
func WrapErrorWithStack(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}

	span := trace.SpanFromContext(ctx)

	uerr := errors.Unwrap(err)
	if uerr == nil {
		uerr = err
	}

	// Record the error with stack trace and set its status in the span.
	span.RecordError(uerr, trace.WithStackTrace(true))
	span.SetStatus(codes.Error, uerr.Error())

	return err
}
