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
	white           = 0
	red             = 31
	green           = 32
	yellow          = 33
	blue            = 34
	violet          = 35
	cyan            = 36
	gray            = 37
)

type StaticFormatter struct {
	CallerPrettyfier func(*runtime.Frame) (function string, file string)
}

func (f *StaticFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	// 写入时间
	time := entry.Time.Format(timestampFormat)
	fmt.Fprintf(b, "%s", time)
	fmt.Fprintf(b, "%s", " ")
	// 写入级别
	levelColor := f.levelColor(entry.Level)
	level := f.level(entry.Level)
	fmt.Fprintf(b, "%c[0;0;%dm%5s%c[0m", 0x1B, levelColor, level, 0x1B)
	fmt.Fprintf(b, "%s", " ")
	// 写入协程 ID
	goroutineID := GetGoroutineID()
	fmt.Fprintf(b, "%c[0;0;%dm%04d%c[0m", 0x1B, violet, goroutineID, 0x1B)
	fmt.Fprintf(b, "%s", " - ")
	// 写入文件和行号
	fileLine := f.fileLine(entry)
	fmt.Fprintf(b, "%c[0;0;%dm%-24s%c[0m", 0x1B, cyan, fileLine, 0x1B)
	fmt.Fprintf(b, "%s", " --- ")
	// 写入消息
	fmt.Fprintf(b, "%s", entry.Message)

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *StaticFormatter) levelColor(level logrus.Level) int {
	switch level {
	case logrus.DebugLevel, logrus.TraceLevel:
		return white
	case logrus.InfoLevel:
		return green
	case logrus.WarnLevel:
		return yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return red
	default:
		return gray
	}
}

func (f *StaticFormatter) level(level logrus.Level) string {
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

func (f *StaticFormatter) fileLine(entry *logrus.Entry) string {
	splits := strings.Split(entry.Caller.File, "/")
	return fmt.Sprintf("%s:%d", splits[len(splits)-1], entry.Caller.Line)
}

func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
