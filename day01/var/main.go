package main

import "fmt"

//声明变量
var a string
var b int

//声明变量（声明全局变量常用）
var (
	name string = "123"
	age  int
	isOk bool
)

func foo() (int, string) {
	return 10, "wtr"
}

func main() {
	//var unuse string	非全局变量声明未使用无法通过编译
	name = "理想"
	age = 16
	isOk = true
	fmt.Printf("name:%s\n", name)
	fmt.Println()
	fmt.Println(age)
	fmt.Println(isOk)
	//声明变量同时赋值
	var title string = "learning golang"
	fmt.Println(title)
	//类型推导
	var lenght = 50
	//一次性声明
	var c, d = "abc", 123
	//简短声明(局部变量声明常用，只能在函数中使用)
	e := 123
	fmt.Println(lenght)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	//匿名变量 用下划线表示
	//匿名变量不占用命名空间 不分配内存 譬如一个函数返回两个变量，但只使用一个
	_, x := foo()
	y, _ := foo()
	fmt.Println(x)
	fmt.Println(y)
}
