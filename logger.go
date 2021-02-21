package sosolog

import (
	"github.com/sirupsen/logrus"
	"os"
)

func New() *logrus.Logger {
	return &logrus.Logger{
		Out:          os.Stdout,
		Level:        logrus.DebugLevel,
		ReportCaller: true,
		Formatter:    &StaticFormatter{},
	}
}
