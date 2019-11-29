package main

import "fmt"

func main() {
	s1 := []int{1, 100, 3}
	fmt.Println(s1)
	s1[1] = 2
	fmt.Println(s1)

	//s1[3] = 100 数组越界

	//append 为切片追加元素 如果切片没有初始化append会自动初始化
	//append的返回值需要变量接收
	s1 = append(s1, 4)
	fmt.Println(s1)

	//append()添加元素和切片扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	var citySlice []string
	// 追加一个元素
	citySlice = append(citySlice, "北京")
	// 追加多个元素
	citySlice = append(citySlice, "上海", "广州", "深圳")
	// 追加切片
	a := []string{"成都", "重庆"}
	citySlice = append(citySlice, a...)
	fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]
}
