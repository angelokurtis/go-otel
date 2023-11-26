package trace

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/sdk/resource"
)

func NewResource(ctx context.Context) (*resource.Resource, error) {
	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithOS(),
		resource.WithContainer(),
		resource.WithHost(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create new resource: %w", err)
	}

	return res, nil
}
