package main

import "fmt"

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	//闭包=函数+引用环境
	//闭包就是一个函数，这个函数包含了它外部作用域的变量

	//底层原理
	//函数可以作为返回值
	//函数内部查找变量先内部找后外部找

	var f = adder()

	//变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。
	//在f的生命周期内，变量x也一直有效。

	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	f1 := adder()
	fmt.Println(f1(40)) //40
	fmt.Println(f1(50)) //90
}
