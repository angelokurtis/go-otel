package otel

import (
	"github.com/google/wire"
	"go.opentelemetry.io/otel/sdk/resource"

	"github.com/angelokurtis/go-starter-otel/internal/env"
	"github.com/angelokurtis/go-starter-otel/internal/trace"
)

var providers = wire.NewSet(
	env.Providers,
	resource.Environment,
	trace.Providers,
	wire.Bind(new(Provider), new(*provider)),
	wire.Struct(new(provider), "*"),
)
