package trace

import (
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	NewSampler,
	NewSpanExporters,
	NewTraceClient,
	NewTracerProvider,
)
