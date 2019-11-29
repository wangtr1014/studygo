package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	//由于返回值没有声明 return x 相当于为返回值开辟内存空间 值为x 后续defer更改x不影响返回值的值
	//1.返回值赋值 2.defer 3.RET指令
	return x //5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	//由于已经声明了返回值是x defer更改了x的值 所以返回值是被修改了
	return 5 //6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	//声明了返回值y y=x=5 defer更改x 不影响y
	return x //5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	//数据类型传参是拷贝 所以defer里修改的x不是返回值x
	return 5 //5
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	//返回值x=5 defer修改拷贝的x不影响返回值
	return 5 //5
}

func f6() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	//返回值x=5 defer修改的是指针类型的引用传递 影响了x的值
	return 5 //6
}

func main() {
	//在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。
	//而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
	fmt.Println(f6())
}
