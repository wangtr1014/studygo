package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

//命名的返回值就相当于在函数中声明一个变量
//return后可以省略 效果通sum
func sum1(x int, y int) (ret int) {
	ret = x + y
	return
}
func func1(x int, y int) {

}

//参数类型相同可简写
func func2(x, y int) {

}
func func3() {

}
func func4() int {
	return 3
}
func func5() (int, int) {
	return 3, 4
}

//参数类型相同可简写
func func6(x, y int, a, b string) {

}

//可变长参数
func func7(x int, y ...string) {
	//y的类型实际是个切片
	//y可传多个也可不传
	//可变长参数必须在函数参数的最后，比如这里不能把y放在前面
}

func main() {
	z := sum(1, 2)
	fmt.Println(z)
}
