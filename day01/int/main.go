package main

import "fmt"

func main() {
	a := 10
	//0开头是八进制
	b := 077
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	//0x开头的是十六进制
	c := 0x123456
	fmt.Printf("%d\n", c)
	//%d十进制 %b二进制 %o八进制 %x十六进制
	//查看变量的类型 %T
	fmt.Printf("%T\n", c)
	//声明一个int8类型的
	d := int8(10)
	fmt.Printf("%T\n", d)
}
