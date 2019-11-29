package main

import (
	"fmt"
	"time"
)

//goroutine类似于线程，属于用户态的线程，我们可以根据需要创建成千上万个goroutine并发工作。
//goroutine是由Go语言的运行时（runtime）调度完成，而线程是由操作系统调度完成
//所以相较于操作系统的线程就比较的轻量级

//Go语言中使用goroutine非常简单，只需要在调用函数的时候在前面加上go关键字，就可以为一个函数创建一个goroutine。
//一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数。

func hello() {
	fmt.Println("hello")
}

func helloInt(i int) {
	fmt.Println(i)
}

//在程序启动时，Go程序就会为main()函数创建一个默认的goroutine，也就是创建一个主goroutine。
func main() {
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main")
	//main函数执行结束了 由main函数启动的goroutine也就都结束了
	//main函数结束的太快 所以不会打印hello

	for i := 0; i < 100; i++ {
		go helloInt(i) //会发现打印出的i是无序的 说明任务是并发异步无序执行的
	}

	//goroutine对应的函数结束 goroutine结束
	//main函数结束 由main函数创建的goroutine也就都结束了

	//利用sleep延缓main函数的退出(这种方式很傻) 即会打印出Hello
	time.Sleep(time.Second)
}
