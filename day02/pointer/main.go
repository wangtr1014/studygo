package main

import "fmt"

func main() {
	//1.&取变量的地址
	n := 10
	p := &n
	fmt.Printf("%T,%v\n", p, p) //*int int类型的指针
	//2.*根据地址取值
	m := *p
	fmt.Printf("%T,%v\n", m, m) //int类型的变量

	// var a *int
	// *a = 100 //只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值
	// fmt.Println(*a)

	//Go语言中new和make是内建的两个函数，主要用来分配内存。
	//使用new函数得到的是一个类型的指针
	var a1 *int
	var a2 = new(int)
	fmt.Println(a1, a2)
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(*a2)
	//make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建
	//而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。

}
