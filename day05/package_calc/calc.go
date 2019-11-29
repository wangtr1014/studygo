package calc

import "fmt"

import initpackage "github.com/wangtr1014/studygo/day05/package_init"

func init() {
	fmt.Println("calc包的init函数")
}

//Add 求两个数之和
func Add(x, y int) (z int) {
	initpackage.Test()
	z = x + y
	return
}
