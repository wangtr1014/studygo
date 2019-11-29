package main

import "fmt"

func test(a []int) {
	a = append(a, 1)
}

func test1(a *[]int) {
	*a = append(*a, 1)
}

func test2(m map[int]int) {
	m[1] = 10
}

func main() {
	var a []int
	var m map[int]int
	a = make([]int, 10, 10)
	m = make(map[int]int, 10)
	test(a)
	fmt.Println(a)
	test1(&a)
	fmt.Println(a)
	test2(m)
	fmt.Println(m)
}
