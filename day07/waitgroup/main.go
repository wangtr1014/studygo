package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f1() {
	//会发现每次得到的r1 r2都是一样的
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}

func f() {
	//用毫秒设置一个时间种子 避免每次得到的随机数一样
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	f1()
	f()

	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}

	wg.Wait() // 等待所有登记的goroutine都结束
}
