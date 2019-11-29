package main

import "fmt"

func main() {
	b1 := true
	//%v 值 %s字符串
	fmt.Printf("%T %v\n", b1, b1)
	//bool 默认值 false
	var b2 bool
	fmt.Println(b2)
}
