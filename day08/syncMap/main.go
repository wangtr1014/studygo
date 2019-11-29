package main

import (
	"fmt"
	"strconv"
	"sync"
)

var sm = sync.Map{}
var m = make(map[string]int)

func nmap() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func smap() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			sm.Store(key, n)
			value, _ := sm.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	//nmap() //fatal error: concurrent map writes
	smap()
}
