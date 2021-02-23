package sosolog

import (
	"testing"
)

func TestNew(t *testing.T) {
	Log := NewLogger()
	Log.Info("测试一下")
	Log.Warn("测试一下")
	Log.Debug("测试一下")
	Log.Error("测试一下")
	Log.Error("测试一下")
}

func TestEvent(t *testing.T) {
	events := []*Event{
		{
			Name:   "WEBRTC",
			Color:  Cyan,
			Hidden: false,
		},
		{
			Name:   "SIGNAL",
			Color:  Green,
			Hidden: false,
		},
	}
	Log := NewEventLogger(events)
	Log.Info("[WEBRTC] 测试 [SIGNAL] 一下")
	Log.Warn("[WEBRTC] 测试一下")
	Log.Debug("[SIGNAL] 测试一下")
	Log.Error("[WEBRTC] 测试一下")
	Log.Error("[SIGNAL] 测试一下")
}
