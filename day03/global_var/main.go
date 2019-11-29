package main

import "fmt"

//通常我们会在if条件判断、for循环、switch语句上使用这种定义变量的方式。

func testLocalVar2(x, y int) {
	fmt.Println(x, y) //函数的参数也是只在本函数中生效
	if x > 0 {
		z := 100 //变量z只在if语句块生效
		fmt.Println(z)
	}
	//fmt.Println(z)//此处无法使用变量z
}

//还有我们之前讲过的for循环语句中定义的变量，也是只在for语句块中生效：

func testLocalVar3() {
	for i := 0; i < 10; i++ {
		fmt.Println(i) //变量i只在当前for语句块中生效
	}
	//fmt.Println(i) //此处无法使用变量i
}

func main() {
	//变量的作用域
	//先再函数内部找
	//其次再函数外面找
	//最后找到全局变量
}
