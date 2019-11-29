package main

import (
	"encoding/json"
	"fmt"
)

//1.序列化 go语言中的结构体变量->json格式的字符串
//2.反序列化 json格式的字符串->go语言中的结构体变量

type company struct {
	name    string
	workers []person
}

type person struct {
	Name string `json:"name"` //指定在json里的名称
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "wtr",
		Age:  27,
	}
	fmt.Printf("%#v\n", p1)
	//json包如果想访问p1下的属性 属性名必须首字母大写
	//序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%#v\n", string(b))
	//反序列化
	str := `{"name":"zhangwei","age":27}`
	var p2 person
	//结构体是值类型 需传指针 否则赋值的只是拷贝
	err = json.Unmarshal([]byte(str), &p2)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%#v\n", p2)
}
