package main

import "fmt"

//recover()必须搭配defer使用。
//defer一定要在可能引发panic的语句之前定义。
//recover很少使用
//连接数据库，defer释放连接,panic 一个应用场景

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println(err)
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
