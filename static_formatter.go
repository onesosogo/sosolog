package sosolog

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
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
	levelColor := levelColor(entry.Level)
	level := levelName(entry.Level)
	fmt.Fprintf(b, "%c[0;0;%dm%5s%c[0m", 0x1B, levelColor, level, 0x1B)
	fmt.Fprintf(b, "%s", " ")
	// 写入协程 ID
	goroutineID := GetGoroutineID()
	fmt.Fprintf(b, "%c[0;0;%dm%05d%c[0m", 0x1B, Violet, goroutineID, 0x1B)
	fmt.Fprintf(b, "%s", " - ")
	// 写入文件和行号
	fileLine := fileLine(entry)
	fmt.Fprintf(b, "%c[0;0;%dm%-14s%c[0m", 0x1B, Cyan, fileLine, 0x1B)
	fmt.Fprintf(b, "%s", " --- ")
	// 写入消息
	fmt.Fprintf(b, "%s", entry.Message)

	b.WriteByte('\n')
	return b.Bytes(), nil
}
