package main

//包名是从$GOPATH/src/后开始计算的，使用/进行路径分隔

import (
	"fmt"
	//import 别名 "包的路径"
	calc "github.com/wangtr1014/studygo/day05/package_calc"
)

//init()函数执行顺序
//Go语言包会从main包开始检查其导入的所有包，每个包中又可能导入了其他的包。
//Go编译器由此构建出一个树状的包引用关系，再根据引用顺序决定编译顺序，依次编译这些包的代码。
//在运行时，被最后导入的包会最先初始化并调用其init()函数

//init函数在全局声明后自动调用，不能主动调用，在main方法执行前执行
func init() {
	fmt.Println("main包的init函数")
}

func main() {
	var (
		x = 1
		y = 2
	)
	calc.Add(x, y)
}
