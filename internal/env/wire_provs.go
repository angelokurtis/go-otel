package env

import (
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	LookupOTel,
	ToTrace,
)
