package log

import (
	"os"
	"github.com/op/go-logging"
)

var normalFormat = `%{color}[PURRGIL] %{time:15:04:05} ▶ %{level} %{color:reset} :: %{message}`
var debugFormat = `%{color}[PURRGIL] (%{time:15:04:05}) %{module} > %{shortfile} ▶ %{level} %{color:reset}:: %{message}`

func ConfigureLoggin(debug bool) {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, getLogFormat(debug))
	backendLeveled := logging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(getLogLevel(debug), "")

	logging.SetBackend(backendLeveled)
}


func getLogFormat(debug bool) logging.Formatter {
	if (debug) {
		return logging.MustStringFormatter(debugFormat)
	}

	return logging.MustStringFormatter(normalFormat)
}

func getLogLevel(debug bool) logging.Level {
	if (debug) {
		return logging.DEBUG
	}

	return logging.INFO
}
