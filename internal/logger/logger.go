package logger

import (
	"log"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
	"go.opentelemetry.io/otel"
)

func New() logr.Logger {
	l := stdr.New(log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile))
	otel.SetLogger(l)

	return l
}
