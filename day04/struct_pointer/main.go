package main

import "fmt"

type people struct {
	name string
	age  int
}

func main() {
	//通过new申请内存的返回的是一个指针
	//这里返回的是这个结构体实例化对应的地址
	p := new(people)
	fmt.Printf("%T\n", p)  //*main.people
	fmt.Printf("%#v\n", p) //&main.people{name:"", age:0}
	//Go语言中支持对结构体指针直接使用.来访问结构体的成员
	p.name = "wtr"
	fmt.Printf("%#v\n", p) //&main.people{name:"wtr", age:0}

	p2 := &people{}         //相当于p2=new(people)
	fmt.Printf("%#v\n", p2) //&main.people{name:"", age:0}
}
