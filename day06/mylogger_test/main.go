package main

import (
	mylogger "github.com/wangtr1014/studygo/day06/mylogger"
)

var log mylogger.Logger

func main() {
	log = mylogger.Newlog("debug")
	s := "这里是具体的错误"
	log.Debug("这是一条格式化后的debug,%s", s)
	log = mylogger.NewFileLog("debug", "./", "file.log", 10*1024)
	log.Error("这是一条格式化后的error,%s", s)
	log.Error("这是一条格式化后的error,%s", s)
}
