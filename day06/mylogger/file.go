package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//向文件里去写日志

//FileLogger 结构体
type FileLogger struct {
	Level       logLevel
	filePath    string //日志文件的路径
	fileName    string //日志文件的文件名
	maxFileSize int64  //文件的最大长度 用于文件切割的判断
	fileObj     *os.File
	errFileObj  *os.File
}

//NewFileLog 实例化一个logger
//结构体比较大 返回指针
func NewFileLog(levelString, fp, fn string, maxSize int64) *FileLogger {
	level := levelParse(levelString)
	f := &FileLogger{
		Level:       level,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	return f
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed %v", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed %v", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return err
}

//Log 判断日志是否打印 打印日志
func (f FileLogger) Log(level logLevel, format string, args ...interface{}) {
	err := f.initFile() //按照文件路径和文件名打开文件
	if err != nil {
		fmt.Printf("open file failed %v", err)
	}
	if f.Level <= level {
		msg := fmt.Sprintf(format, args...)
		fileName, funcName, line := getInfo(3)
		now := time.Now()
		timeStr := now.Format("2006-01-02 15:04:05")
		fmt.Fprintf(f.fileObj, "[%s]-%s文件-%s函数-第%d行-[%s]\n", timeStr, fileName, funcName, line, msg)
		if level >= ERROR {
			fmt.Fprintf(f.errFileObj, "[%s]-%s文件-%s函数-第%d行-[%s]\n", timeStr, fileName, funcName, line, msg)
		}
	}
	f.close()
}

func (f FileLogger) close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

//Debug 记录debug日志
//带有格式化的
func (f FileLogger) Debug(format string, args ...interface{}) {
	f.Log(DEBUG, format, args...)
}

//Trace 记录trace日志
func (f FileLogger) Trace(format string, args ...interface{}) {
	f.Log(TRACE, format, args...)
}

//Info 记录info日志
func (f FileLogger) Info(format string, args ...interface{}) {
	f.Log(INFO, format, args...)
}

//Warning 记录warning日志
func (f FileLogger) Warning(format string, args ...interface{}) {
	f.Log(WARNING, format, args...)
}

//Error 记录error日志
func (f FileLogger) Error(format string, args ...interface{}) {
	f.Log(ERROR, format, args...)
}

//Fatal 记录fatal日志
func (f FileLogger) Fatal(format string, args ...interface{}) {
	f.Log(FATAL, format, args...)
}
