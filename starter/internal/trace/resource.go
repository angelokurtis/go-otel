package trace

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/sdk"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func NewResource(ctx context.Context) (*resource.Resource, error) {
	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithOS(),
		resource.WithContainer(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.TelemetrySDKName("opentelemetry"),
			semconv.TelemetrySDKLanguageGo,
			semconv.TelemetrySDKVersion(sdk.Version()),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create new resource: %w", err)
	}

	return res, nil
}
