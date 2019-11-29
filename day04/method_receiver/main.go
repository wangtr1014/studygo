package main

import "fmt"

//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法
//这里的p Person叫做接收者 Dream叫做方法
//函数不属于任何类型，方法属于特定的类型
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// SetAge2 设置p的年龄
// 使用值接收者
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("小王子", 25)
	p1.Dream()
	fmt.Println(p1.age) // 25
	p1.SetAge(30)       // (*p1).SetAge2(30)
	fmt.Println(p1.age) // 30

	p2 := NewPerson("小王子", 25)
	p2.Dream()
	fmt.Println(p2.age) // 25
	p2.SetAge2(30)      // (*p1).SetAge2(30)
	fmt.Println(p2.age) // 25

	// 什么时候应该使用指针类型接收者
	// 需要修改接收者中的值
	// 接收者是拷贝代价比较大的大对象
	// 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
}
