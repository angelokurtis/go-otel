package trace

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type ServiceNameProvider interface {
	ServiceName() (string, bool)
}

func NewResource(ctx context.Context, provider ServiceNameProvider) (*resource.Resource, error) {
	attrs := []attribute.KeyValue{
		semconv.TelemetrySDKName("opentelemetry"),
		semconv.TelemetrySDKLanguageGo,
		semconv.TelemetrySDKVersion(sdk.Version()),
	}

	if serviceName, ok := provider.ServiceName(); ok {
		attrs = append(attrs, semconv.ServiceName(serviceName))
	}

	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithOS(),
		resource.WithContainer(),
		resource.WithHost(),
		resource.WithAttributes(attrs...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create new resource: %w", err)
	}

	return res, nil
}
