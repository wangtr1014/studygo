package main

import (
	"fmt"
	"os"
	"runtime"
)

func run() {
	//层数 n=1 向上一层 调用的行数是26 调用函数是 main.run
	n := 1
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("调用runtime caller失败")
		os.Exit(1)
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	//调用caller函数的文件
	fmt.Println(file)
	//调用caller函数的行数
	fmt.Println(line)
}

func main() {
	run()
	//层数
	n := 0
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("调用runtime caller失败")
		os.Exit(1)
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	//调用caller函数的文件
	fmt.Println(file)
	//调用caller函数的行数
	fmt.Println(line)
}
