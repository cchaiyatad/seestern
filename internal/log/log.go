package log

import (
	l "log"
	"os"
)

type Priority uint8

const (
	Info Priority = iota
	Warning
	Error
)

var logI *l.Logger
var logW *l.Logger
var logE *l.Logger

func init() {
	logI = l.New(os.Stderr, "[    Info]: ", l.Ldate|l.Ltime)
	logW = l.New(os.Stderr, "[ Warning]: ", l.Ldate|l.Ltime)
	logE = l.New(os.Stderr, "[   Error]: ", l.Ldate|l.Ltime)
}

func Logf(priority Priority, format string, v ...interface{}) {
	logger := logHelper(priority)
	logger.Printf(format, v...)
}

func Log(priority Priority, v ...interface{}) {
	logger := logHelper(priority)
	logger.Println(v...)
}

func logHelper(priority Priority) *l.Logger {
	switch priority {
	case Info:
		return logI
	case Warning:
		return logW
	case Error:
		return logE
	default:
		return logE
	}
}
