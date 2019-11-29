package main

import "fmt"

func f1() {

}

func f2() int {
	return 1
}

func f4(x, y int) {

}

//函数也可以作为参数
func f3(x func() int) {
	fmt.Println(x())
}

func ff(y int) (x int) {
	x = y
	fmt.Println(x)
	return x
}

//函数也可以作为返回值
func f5(x func() int) func(int) int {
	return ff
}

func main() {
	a := f1
	fmt.Printf("%T\n", a) //func()
	b := f2
	fmt.Printf("%T\n", b) //func() int
	c := f4
	fmt.Printf("%T\n", c) //func(int, int)
	f3(f2)
	f3(b)
	f5(f2)(10)
}
