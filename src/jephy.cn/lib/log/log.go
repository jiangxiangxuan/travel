package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

var MAX_BYTES_PER_LOG = 1024 * 3

const TimeFormat = "2006/01/02 15:04:05"
const FieldSplit = " - "

const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelBuss
)

var LogLevelString = map[string]int{
	"trace": LevelTrace,
	"debug": LevelDebug,
	"info":  LevelInfo,
	"warn":  LevelWarn,
	"error": LevelError,
	"fatal": LevelFatal,
	"buss":  LevelBuss,
}

var LevelName [7]string = [7]string{
	"TRACE", "DEBUG", "INFO", "WARN",
	"ERROR", "FATAL", "BUSS",
}

type Logger struct {
	level int
	w     io.Writer
}

func (l *Logger) SetLevel(level int) {
	l.level = level
}

func (l *Logger) Output(callDepth int, level int, format string, v ...interface{}) {
	if l.level > level {
		return
	}

	var file_line, now, slevel, msg string

	now = time.Now().Format(TimeFormat)
	slevel = LevelName[level]

	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		file = "???"
		line = 0
	} else {
		v := strings.Split(file, "/")
		idx := len(v) - 3
		if idx < 0 {
			idx = 0
		}
		file = strings.Join(v[idx:], "/")
	}
	file_line = fmt.Sprintf("%s:[%d]", file, line)

	msg = fmt.Sprintf(format, v...)
	if len(msg) > MAX_BYTES_PER_LOG {
		msg = fmt.Sprintf("%s... data too long, soucre-length=%d",
			msg[0:MAX_BYTES_PER_LOG], len(msg))
	}

	// 行号符提成 "||"
	msg = strings.Replace(msg, "\n", "||", -1)
	// 回车符替换成 "||"
	msg = strings.Replace(msg, "\r", "||", -1)

	result := fmt.Sprintf("%s - %s - %s - %s", now, slevel, file_line, msg)
	l.w.Write([]byte(result))
	if result[len(result)-1] != '\n' {
		l.w.Write([]byte("\n"))
	}
}

func NewLogger() *Logger {
	var l = new(Logger)
	l.level = LevelInfo
	l.w = os.Stdout
	return l
}

var std = NewLogger()

func SetLevelS(level string) {
	SetLevel(LogLevelString[strings.ToLower(level)])
}

func SetLevel(level int) {
	std.SetLevel(level)
}

func Trace(format string, v ...interface{}) {
	std.Output(2, LevelTrace, format, v...)
}

func Debug(format string, v ...interface{}) {
	std.Output(2, LevelDebug, format, v...)
}

func Info(format string, v ...interface{}) {
	std.Output(2, LevelInfo, format, v...)
}

func Warn(format string, v ...interface{}) {
	std.Output(2, LevelWarn, format, v...)
}

func Error(format string, v ...interface{}) {
	std.Output(2, LevelError, format, v...)
}

func Fatal(format string, v ...interface{}) {
	std.Output(2, LevelFatal, format, v...)
}

func Buss(format string, v ...interface{}) {
	std.Output(2, LevelBuss, format, v...)
}
