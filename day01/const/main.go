package main

const pi = 3.14

//批量声明常量
const (
	successCode = 1
	failCode    = -1
)

//没有表达式表示和上一行表达式保存一致
const (
	n1 = 100
	n2
	n3
)

//iota是go语言的常量计数器，只能在常量的表达式中使用。
//iota在const关键字出现时将被重置为0。
//const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。
//使用iota能简化定义，在定义枚举时很有用。

const (
	a1 = iota //0
	a2        //1
	a3        //2
)

const (
	b1 = iota //0
	b2        //1
	_
	b4 //3
)

const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)
const n5 = iota //0

const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

func main() {

}
