package sosolog

import (
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger() *logrus.Logger {
	return &logrus.Logger{
		Out:          os.Stdout,
		Level:        logrus.DebugLevel,
		ReportCaller: true,
		Formatter:    &StaticFormatter{},
	}
}

func NewStaticLogger() *logrus.Logger {
	return &logrus.Logger{
		Out:          os.Stdout,
		Level:        logrus.DebugLevel,
		ReportCaller: true,
		Formatter:    &StaticFormatter{},
	}
}

func NewEventLogger(events []*Event) *logrus.Logger {
	return &logrus.Logger{
		Out:          os.Stdout,
		Level:        logrus.DebugLevel,
		ReportCaller: true,
		Formatter: &EventFormatter{
			Events: events,
		},
	}
}
