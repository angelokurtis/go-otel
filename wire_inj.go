//go:build wireinject
// +build wireinject

package otel

import (
	"context"

	"github.com/google/wire"
)

// Init is an initializer for an OpenTelemetry provider that offers a high level of configurability
func Init(ctx context.Context) (Provider, error) {
	wire.Build(providers)
	return nil, nil
}
