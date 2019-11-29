package main

import (
	"fmt"
	"sync"
)

var sum = 0

var wg sync.WaitGroup
var mt sync.Mutex //互斥锁

func test() {
	defer wg.Done()
	for i := 0; i < 50000; i++ {
		mt.Lock()
		sum++
		mt.Unlock()
	}
}

func main() {
	wg.Add(2)
	go test()
	go test()
	wg.Wait()
	fmt.Println(sum)
}
