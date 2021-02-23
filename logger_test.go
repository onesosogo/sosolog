package sosolog

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNew(t *testing.T) {
	Log := New()
	Log.Info("测试一下")
	Log.Warn("测试一下")
	Log.Debug("测试一下")
	Log.Error("测试一下")
	logrus.Info("DDD")
}
