package main

import "fmt"

//接口是一种类型
//规定了变量有哪些方法

//经常遇到以下场景，不关心属性（数据），只关心行为（方法）。
//比如可以创建一个接口进行数据库连接，不关系使用的是哪种数据库
//例如fmt.Print()可以打印多种数据类型

type cat struct {
	name string
}
type dog struct {
	name string
}

func (c cat) speak() {
	fmt.Printf("%s学猫叫\n", c.name)
}
func (d dog) speak() {
	fmt.Printf("%s学狗叫\n", d.name)
}

//定义一个包含speak方法的接口
type speaker interface {
	speak()
}

//如果没有接口类型，得为cat和dog各写一个hit方法
func hit(s speaker) {
	s.speak()
}

func main() {
	c := cat{
		name: "tom",
	}
	d := dog{
		name: "snoopy",
	}
	hit(c)
	hit(d)
}
