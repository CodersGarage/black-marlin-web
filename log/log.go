package log

import (
	"github.com/sirupsen/logrus"
	"os"

	"github.com/CodersGarage/black-marlin-web/log/hooks"
)

var defLogger *logrus.Logger

func Init() {
	defLogger = logrus.New()
	defLogger.Out = os.Stdout
	defLogger.AddHook(hooks.NewHook())
}

func Logger() *logrus.Logger {
	return defLogger
}
