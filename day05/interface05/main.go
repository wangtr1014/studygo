package main

import "fmt"

//空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
//空接口类型的变量可以存储任意类型的变量。

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	// 定义一个空接口x
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)

	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)

	//一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值。
	// 想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：
	// x.(T)
	// x：表示类型为interface{}的变量
	// T：表示断言x可能是的类型。
	// 该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。
	var z interface{}
	z = "Hello 沙河"
	v, ok := z.(int)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}

	func(x interface{}) {
		switch v := x.(type) {
		case string:
			fmt.Printf("x is a string，value is %v\n", v)
		case int:
			fmt.Printf("x is a int is %v\n", v)
		case bool:
			fmt.Printf("x is a bool is %v\n", v)
		default:
			fmt.Println("unsupport type！")
		}
	}(z)
}
