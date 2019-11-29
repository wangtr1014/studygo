package main

import "fmt"

func main() {
	//Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
	//在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行
	//也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
	fmt.Println("strat")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
	//strat
	// end
	// 3
	// 2
	// 1
	//由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。
	//比如：资源清理、文件关闭、解锁及记录时间等。
}
