package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var lock sync.Mutex
var wg sync.WaitGroup

var (
	x1 int64
	x2 int64
	x3 int64
)

func add1() {
	defer wg.Done()
	x1++
}

func add2() {
	defer wg.Done()
	lock.Lock()
	x2++
	lock.Unlock()
}

func add3() {
	defer wg.Done()
	atomic.AddInt64(&x3, 1)
}

func main() {
	wg.Add(150000)
	for i := 0; i < 50000; i++ {
		go add1()
	}
	for i := 0; i < 50000; i++ {
		go add2()
	}
	for i := 0; i < 50000; i++ {
		go add3()
	}
	wg.Wait()
	fmt.Println(x1, x2, x3)
}
