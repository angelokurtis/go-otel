package trace

import (
	xraypropagator "go.opentelemetry.io/contrib/propagators/aws/xray"
	b3propagator "go.opentelemetry.io/contrib/propagators/b3"
	jaegerpropagator "go.opentelemetry.io/contrib/propagators/jaeger"
	otpropagator "go.opentelemetry.io/contrib/propagators/ot"
	"go.opentelemetry.io/otel/propagation"
)

func NewTextMapPropagator(propagators Propagators) propagation.TextMapPropagator {
	props := make([]propagation.TextMapPropagator, 0, len(propagators))

	for _, propagator := range propagators {
		switch propagator {
		case PropagatorTracecontext:
			props = append(props, propagation.TraceContext{})
		case PropagatorBaggage:
			props = append(props, propagation.Baggage{})
		case PropagatorB3:
			props = append(props, b3propagator.New(b3propagator.WithInjectEncoding(b3propagator.B3SingleHeader)))
		case PropagatorB3multi:
			props = append(props, b3propagator.New(b3propagator.WithInjectEncoding(b3propagator.B3MultipleHeader)))
		case PropagatorJaeger:
			props = append(props, jaegerpropagator.Jaeger{})
		case PropagatorXray:
			props = append(props, xraypropagator.Propagator{})
		case PropagatorOttrace:
			props = append(props, otpropagator.OT{})
		}
	}

	return propagation.NewCompositeTextMapPropagator(props...)
}
