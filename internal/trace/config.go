package trace

import (
	"net/url"
	"time"
)

type Config struct {
	Exporters   []Exporter
	Endpoint    url.URL
	Timeout     time.Duration
	Protocol    Protocol
	Compression Compression
	Sampler     Sampler
	SamplerArg  float64
	Propagators []Propagator
}
