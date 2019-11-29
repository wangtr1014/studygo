package main

import "fmt"

func isMan(age int) {
	if age > 159 {
		fmt.Println("成仙了")
	} else if age >= 18 {
		fmt.Println("成年了")
	} else {
		fmt.Println("未成年")
	}
	if in := 8; in < 10 {
		fmt.Println("太小了")
	}
}

func main() {
	age := 19
	isMan(age)
	//循环
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("---------------------")
	//变种
	i := 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("---------------------")
	//变种
	i = 10
	for i < 15 {
		fmt.Println(i)
		i++
	}
	//无限循环
	// for {
	// 	fmt.Println("go on")
	// }
	//for循环可以通过break、goto、return、panic语句强制退出循环。

	// for range(键值循环)
	// Go语言中可以使用for range遍历数组、切片、字符串、map 及通道（channel）。 通过for range遍历的返回值有以下规律：
	// 数组、切片、字符串返回索引和值。
	// map返回键和值。
	// 通道（channel）只返回通道内的值。

	str := "hello world"
	for i, v := range str {
		fmt.Printf("%d-%c\n", i, v)
	}

	str1 := "hello大家好"
	for i, v := range str1 {
		fmt.Printf("%d-%c\n", i, v)
	}
	for i, v := range str1 {
		fmt.Printf("%d-%c\n", i, v)
	}

	str2 := []byte(str1)
	for i, v := range str2 {
		fmt.Printf("%d-%c\n", i, v)
	}

	str3 := []rune(str1)
	for i, v := range str3 {
		fmt.Printf("%d-%c\n", i, v)
	}
}
