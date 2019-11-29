package main

import "fmt"

//结构体模拟其它语言中的继承

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Println("可以动弹")
}

type dog struct {
	feet uint8
	//这样dog结构体实例化后就可以使用animal结构体的属性和方法
	animal //匿名结构体 方便直接使用dog.name
}

func (d dog) shout() {
	fmt.Printf("%s在叫唤\n", d.name)
}

func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "dog",
		},
	}
	d1.shout()
	d1.move()
}
