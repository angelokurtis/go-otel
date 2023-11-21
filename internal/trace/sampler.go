package trace

import (
	"fmt"

	"go.opentelemetry.io/otel/sdk/trace"
)

type SamplerOptions struct {
	Sampler    Sampler
	SamplerArg SamplerArg
}

func NewSampler(options SamplerOptions) (trace.Sampler, error) {
	switch sampler := options.Sampler; sampler {
	case SamplerAlwaysOn:
		return trace.AlwaysSample(), nil
	case SamplerAlwaysOff:
		return trace.NeverSample(), nil
	case SamplerTraceidratio:
		fraction := float64(options.SamplerArg)
		return trace.TraceIDRatioBased(fraction), nil
	case SamplerParentbasedAlwaysOn:
		return trace.ParentBased(trace.AlwaysSample()), nil
	case SamplerParentbasedAlwaysOff:
		return trace.ParentBased(trace.NeverSample()), nil
	case SamplerParentbasedTraceidratio:
		fraction := float64(options.SamplerArg)
		return trace.ParentBased(trace.TraceIDRatioBased(fraction)), nil
	default:
		return nil, fmt.Errorf("unsupported sampler option: %s", sampler)
	}
}
