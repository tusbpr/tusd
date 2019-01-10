package cli

import (
	"log"
	"os"

	"github.com/tus/tusd"
)

var stdout = log.New(os.Stdout, "[tusd] ", 0)
var stderr = log.New(os.Stderr, "[tusd] ", 0)

func logEv(logOutput *log.Logger, eventName string, details ...string) {
	tusd.LogEvent(logOutput, eventName, details...)
}
