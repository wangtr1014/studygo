package main

import "fmt"

func main() {
	//切片（Slice）是一个拥有相同类型元素的可变长度的序列。
	//它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
	//切片是一个引用类型(数组是值类型)，它的内部结构包含地址、长度和容量
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) //nil相当于空
	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"北京", "上海", "广州", "深圳"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) //nil相当于空
	//长度和容量 len cap
	fmt.Printf("长度:%d,容量:%d\n", len(s1), cap(s1))
	//由数组得到切片
	s3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	s4 := s3[0:3] //基于一个数组进行切割，左包含，右不包含
	fmt.Println(s4)
	fmt.Printf("长度:%d,容量:%d\n", len(s4), cap(s4))
	s4 = s3[:3]
	fmt.Println(s4)
	fmt.Printf("长度:%d,容量:%d\n", len(s4), cap(s4))
	s4 = s3[2:]
	fmt.Println(s4)
	fmt.Printf("长度:%d,容量:%d\n", len(s4), cap(s4))
	s4 = s3[:]
	fmt.Println(s4)
	fmt.Printf("长度:%d,容量:%d\n", len(s4), cap(s4))
	//切片的容量是指底层数组的融科了
	//底层数组从切片的第一个元素到最后的元素的数量

	//切片是引用类型，指向底层的一个具体数组，数据的值变了，切片的值也对应变化
	a := [...]int{1, 2, 3, 4, 5, 6}
	q := a[0:3]
	fmt.Println(a)
	fmt.Println(q)
	a[2] = 1000
	fmt.Println(a)
	fmt.Println(q)
	fmt.Printf("长度:%d,容量:%d\n", len(q), cap(q))
	//切片可以再由切片分割得到，共享底层数组
}
