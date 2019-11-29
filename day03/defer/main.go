package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	//defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	//先转换成 defer calc("AA", 1, 3)
	x = 10
	defer calc("BB", x, calc("B", x, y))
	//先转换成 defer calc("BB", 10, 12)
	y = 20
	//BB 10 12 22
	//AA 1 3 4
}
