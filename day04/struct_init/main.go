package main

import "fmt"

type person struct {
	name string
	age  int
}

//struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大
//所以该构造函数返回的是结构体指针类型。
func newPersonPointer(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func main() {
	//没有初始化的结构体，其成员变量都是对应其类型的零值。
	p1 := person{
		name: "小王子",
		age:  18,
	}
	fmt.Printf("%#v\n", p1)
	p2 := &person{
		name: "小王子",
		age:  18,
	}
	fmt.Printf("%#v\n", p2)
	p3 := &person{
		name: "小王子",
	}
	fmt.Printf("%#v\n", p3)
	/*
		    必须初始化结构体的所有字段。
		    初始值的填充顺序必须与字段在结构体中的声明顺序一致。
			该方式不能和键值初始化方式混用。
	*/
	p4 := &person{
		"小王子",
		18,
	}
	fmt.Printf("%#v\n", p4)

	//构造函数
	p5 := newPerson("wtr", 25)
	fmt.Printf("%#v\n", p5)

	// 结构体占用一块连续的内存。
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
}
