package logger

import (
	"log"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
)

func New() logr.Logger {
	l := stdr.New(log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile))

	return l
}
