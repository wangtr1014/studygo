package main

import "fmt"

func main() {
	//值类型为map的切片
	s := make([]map[int]string, 0, 10)
	//s[0][0] = "first" index out of range make初始化了切片，没有初始化map
	m := make(map[int]string, 1)
	s = append(s, m)
	s[0][0] = "first"
	fmt.Println(s)

	//值为切片类型的map
	m1 := make(map[string][]int, 10)
	m1["test"] = []int{1, 2, 3}
	fmt.Println(m1)
}
