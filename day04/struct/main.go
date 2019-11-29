package main

import "fmt"

type person struct {
	name  string
	age   int
	city  string
	hobby []string
}

type people struct {
	name string
	age  int
}

//匿名结构体	多用于临时场景
var s struct {
	name string
}

func main() {
	// 	只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
	// 结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。
	var p person
	p.name = "Jerry"
	p.hobby = []string{"Ball", "Music"}
	fmt.Println(p)

	//结构体是值类型
	var peo people
	peo.name = "wtr"
	peo.age = 18

	func(x people) {
		//更改的是一个副本
		x.age = 30
	}(peo)
	fmt.Println(peo.age) //18

	func(x *people) {
		//更改的是同一块内存
		(*x).age = 30
	}(&peo)
	fmt.Println(peo.age) //30
	fmt.Printf("%v\n", peo)
	fmt.Printf("%#v\n", peo)
}
