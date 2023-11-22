package otel

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func WrapError(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}

	span := trace.SpanFromContext(ctx)

	uerr := errors.Unwrap(err)
	if uerr == nil {
		uerr = err
	}

	span.RecordError(uerr)
	span.SetStatus(codes.Error, uerr.Error())

	return err
}
