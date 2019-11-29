package main

import (
	"fmt"
	"sort"
)

func main() {
	//Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中

	// copy()复制切片
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	//var c []int 这么写不能copy成功，因为c没有被分配内存空间
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
	c[0] = 1000
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1000 2 3 4 5]

	//Go语言中并没有删除切片元素的专用方法
	// 从切片中删除元素
	b := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	b = append(b[:2], b[3:]...)
	fmt.Println(b) //[30 31 33 34 35 36 37]
	//要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)

	var d = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		d = append(d, i)
	}
	fmt.Println(d)

	var n = [...]int{1, 9, 2, 5, 3, 15}
	fmt.Println(n)
	s := n[:]
	sort.Ints(s) //排序
	fmt.Println(s)
	fmt.Println(n) //排序切片后数组也发生了变化

	x := [...]int{1, 3, 5, 7, 9, 13, 15}
	fmt.Println(x)
	y := append(x[0:1], x[2:]...)
	fmt.Println(y)
	fmt.Println(x) //切片变化实际是底层数组变化

}
