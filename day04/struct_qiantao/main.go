package main

import "fmt"

type address struct {
	province string
	city     string
}

type company struct {
	name string
	addr address
}

// type company struct {
// 	name string
// 	address 匿名嵌套结构体
// }

func main() {
	c := company{
		name: "Huoli",
		addr: address{
			province: "安徽",
			city:     "合肥",
		},
	}
	fmt.Printf("%#v\n", c)
	fmt.Println(c.name)
	fmt.Println(c.addr.city)
	//fmt.Println(c.city) 如果是使用匿名嵌套结构体是可以直接取city的
	//先在自身结构体查找city 找不到会自动到匿名结构体里找
	//如果有多个匿名嵌套结构体 都包含city属性 则不能这么取city的值
}
