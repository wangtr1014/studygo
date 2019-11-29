package main

import (
	"fmt"
	"math"
)

func main() {
	f1 := math.MaxFloat32
	fmt.Printf("最大的32位浮点数：%f\n", f1)
	//默认是64位
	f2 := 1.23456
	fmt.Printf("%T\n", f2)
	f3 := float32(1.23456)
	fmt.Printf("%T\n", f3)
	//f2=f3 报错 类型不同
}
