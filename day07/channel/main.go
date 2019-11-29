package main

import "fmt"

//channel是一种类型，一种引用类型。声明通道类型的格式如下：
//var 变量 chan 元素类型

var b chan int

func main() {
	fmt.Println(b) //nil
	//通道必须初始化后才能使用
	//make 为slice map chan初始化
	b = make(chan int) //(必须有人接收才能放到通道里，否则等待)

	// b <- 10 如果直接执行就卡主了，程序报错 因为没有接收者

	b = make(chan int, 16) //带缓冲区的通道（无人接收最多放16个数据，再多只能等待）
	fmt.Println(b)

	//通道有发送（send）、接收(receive）和关闭（close）三种操作
	//ch <- 10 把10发送到ch中
	//x := <- ch 从ch中接收值并赋值给变量x
	//<-ch 从ch中接收值，忽略结果
	//close(ch)我们通过调用内置的close函数来关闭通道。
}
