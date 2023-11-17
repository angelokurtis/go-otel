package trace

import (
	"context"

	xraypropagator "go.opentelemetry.io/contrib/propagators/aws/xray"
	b3propagator "go.opentelemetry.io/contrib/propagators/b3"
	jaegerpropagator "go.opentelemetry.io/contrib/propagators/jaeger"
	otpropagator "go.opentelemetry.io/contrib/propagators/ot"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func NewTextMapPropagator(ctx context.Context, config *Config) propagation.TextMapPropagator {
	propagators := make([]propagation.TextMapPropagator, 0, len(config.Propagators))

	for _, propagator := range config.Propagators {
		switch propagator {
		case PropagatorTracecontext:
			propagators = append(propagators, propagation.TraceContext{})
		case PropagatorBaggage:
			propagators = append(propagators, propagation.Baggage{})
		case PropagatorB3:
			propagators = append(propagators, b3propagator.New(b3propagator.WithInjectEncoding(b3propagator.B3SingleHeader)))
		case PropagatorB3multi:
			propagators = append(propagators, b3propagator.New(b3propagator.WithInjectEncoding(b3propagator.B3MultipleHeader)))
		case PropagatorJaeger:
			propagators = append(propagators, jaegerpropagator.Jaeger{})
		case PropagatorXray:
			propagators = append(propagators, xraypropagator.Propagator{})
		case PropagatorOttrace:
			propagators = append(propagators, otpropagator.OT{})
		}
	}

	propagator := propagation.NewCompositeTextMapPropagator(propagators...)
	otel.SetTextMapPropagator(propagator)

	return propagator
}
