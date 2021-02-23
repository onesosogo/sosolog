package sosolog

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strconv"
	"strings"
)

const (
	logFormat       = "%time% %lvl% %gid% %position% --- %msg%"
	timestampFormat = "2006-01-02 15:04:05.000"
)

type Color int

const (
	White  Color = 0
	Red          = 31
	Green        = 32
	Yellow       = 33
	Blue         = 34
	Violet       = 35
	Cyan         = 36
	Gray         = 37
)

func levelColor(level logrus.Level) Color {
	switch level {
	case logrus.DebugLevel, logrus.TraceLevel:
		return White
	case logrus.InfoLevel:
		return Green
	case logrus.WarnLevel:
		return Yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return Red
	default:
		return Gray
	}
}

func levelName(level logrus.Level) string {
	switch level {
	case logrus.DebugLevel:
		return "DEBUG"
	case logrus.TraceLevel:
		return "TRACE"
	case logrus.InfoLevel:
		return "INFO"
	case logrus.WarnLevel:
		return "WARN"
	case logrus.ErrorLevel:
		return "ERROR"
	case logrus.FatalLevel:
		return "FATAL"
	case logrus.PanicLevel:
		return "PANIC"
	default:
		return "UNKNOWN"
	}
}

func fileLine(entry *logrus.Entry) string {
	splits := strings.Split(entry.Caller.File, "/")
	file := splits[len(splits)-1]
	files := strings.Split(file, ".")
	fileName := files[0]
	fileAndLine := fmt.Sprintf("%s:%d", fileName, entry.Caller.Line)
	count := len(fileAndLine)
	if count <= 14 {
		return fileAndLine
	} else {
		pre := ".."
		rs := []rune(fileAndLine)
		end := count
		begin := count - 10
		return pre + string(rs[begin:end])
	}
}

func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func eventColorMessage(events []*Event, message string) string {
	for _, e := range events {
		if e.Name != "" {
			oldPattern := "[" + e.Name + "]"
			newPattern := fmt.Sprintf("%c[0;0;%dm%s%c[0m", 0x1B, e.Color, oldPattern, 0x1B)
			message = strings.Replace(message, oldPattern, newPattern, -1)
		}
	}
	return message
}

func checkEvent(events []*Event, message string) bool {
	for _, e := range events {
		if e.Name != "" {
			pattern := "[" + e.Name + "]"
			if strings.Contains(message, pattern) && e.Hidden {
				return true
			}
		}
	}
	return false
}
