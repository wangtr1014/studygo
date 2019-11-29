package mylogger

import (
	"fmt"
	"time"
)

//向终端去写日志

//ConsoleLogger 结构体
type ConsoleLogger struct {
	Level logLevel
}

//Newlog 实例化一个logger
func Newlog(levelString string) ConsoleLogger {
	level := levelParse(levelString)
	return ConsoleLogger{
		Level: level,
	}
}

//Log 判断日志是否打印 打印日志
func (l ConsoleLogger) Log(level logLevel, format string, args ...interface{}) {
	if l.Level <= level {
		msg := fmt.Sprintf(format, args...)
		fileName, funcName, line := getInfo(3)
		now := time.Now()
		timeStr := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s]-%s文件-%s函数-第%d行-[%s]\n", timeStr, fileName, funcName, line, msg)
	}
}

//Debug 记录debug日志
//带有格式化的
func (l ConsoleLogger) Debug(format string, args ...interface{}) {
	l.Log(DEBUG, format, args...)
}

//Trace 记录trace日志
func (l ConsoleLogger) Trace(format string, args ...interface{}) {
	l.Log(TRACE, format, args...)
}

//Info 记录info日志
func (l ConsoleLogger) Info(format string, args ...interface{}) {
	l.Log(INFO, format, args...)
}

//Warning 记录warning日志
func (l ConsoleLogger) Warning(format string, args ...interface{}) {
	l.Log(WARNING, format, args...)
}

//Error 记录error日志
func (l ConsoleLogger) Error(format string, args ...interface{}) {
	l.Log(ERROR, format, args...)
}

//Fatal 记录fatal日志
func (l ConsoleLogger) Fatal(format string, args ...interface{}) {
	l.Log(FATAL, format, args...)
}
