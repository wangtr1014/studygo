package mylogger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

type logLevel uint8

const (
	DEBUG   logLevel = 1
	TRACE   logLevel = 2
	INFO    logLevel = 3
	WARNING logLevel = 4
	ERROR   logLevel = 5
	FATAL   logLevel = 6
)

type Logger interface {
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}

func getInfo(n int) (fileName, funcName string, line int) {
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("调用runtime caller失败")
		os.Exit(1)
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return fileName, funcName, line
}

func levelParse(str string) (level logLevel) {
	str = strings.ToLower(str)
	switch str {
	case "debug":
		level = DEBUG
	case "trace":
		level = TRACE
	case "info":
		level = INFO
	case "warning":
		level = WARNING
	case "error":
		level = ERROR
	case "fatal":
		level = FATAL
	default:
		level = FATAL
	}
	return
}
