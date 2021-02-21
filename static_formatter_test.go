package sosolog

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
	"time"
)

func TestStaticFormatter_Format(t *testing.T) {
	log := &logrus.Logger{
		Out:          os.Stdout,
		Level:        logrus.DebugLevel,
		ReportCaller: true,
		Formatter:    &StaticFormatter{},
	}
	log.Debug("测试内容", 30, false)
	log.Info("测试内容")
	log.Warn("测试内容")
	log.Error("测试内容")
	for i := 0; i < 10000; i++ {
		go func() {
			log.Info("测试内容")
		}()
	}
	time.Sleep(time.Minute)
}
