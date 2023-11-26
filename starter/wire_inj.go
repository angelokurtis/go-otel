//go:build wireinject
// +build wireinject

package otel

import (
	"context"

	"github.com/google/wire"
)

func NewManager(ctx context.Context) (*Manager, func(), error) {
	wire.Build(providers)
	return nil, nil, nil
}
