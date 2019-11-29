package main

import "fmt"

func main() {
	//数组 存放元素的容器
	//指定存放元素的类型和容量（长度）
	var a1 [3]bool //[true false true]
	var a2 [4]bool //[true true false fasle]
	fmt.Printf("a1:%T,a2:%T\n", a1, a2)
	//长度是数组类型的一部分，a1 a2不能进行比较
	//a=b会报错 在golang里切片的应用比较多

	//数组的初始化
	//如果未初始化，默认元素都是零值(也就是默认值false 0 "")
	//method1
	a3 := [3]bool{true, true, true}
	fmt.Println(a3)
	//method2 根据初始值自动推断长度
	a4 := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(a4)
	//method3
	a5 := [5]int{1, 2}
	fmt.Println(a5)
	a6 := [5]int{0: 1, 4: 2} //根据索引
	fmt.Println(a6)

	//数组的遍历
	citys := [...]string{"北京", "深圳", "武汉", "合肥"}
	//method1 根据索引
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//method2 range
	for index, value := range citys {
		fmt.Println(index, value)
	}

	//多维数组
	var aa [3][2]int
	aa = [3][2]int{
		[2]int{1, 2},
		[2]int{2, 3},
		[2]int{3, 4}, //这里居然必须有逗号
	}
	fmt.Println(aa)
	for _, v1 := range aa {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
