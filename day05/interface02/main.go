package main

import "fmt"

// type 接口类型名 interface{
//     方法名1( 参数列表1 ) 返回值列表1
//     方法名2( 参数列表2 ) 返回值列表2
//     …
// }

//一个对象只要全部实现了接口中的方法，那么就实现了这个接口。
//换句话说，接口就是一个需要实现的方法列表。

//接口名：使用type将接口定义为自定义的类型名。
//Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，
//有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
//方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，
//这个方法可以被接口所在的包（package）之外的代码访问。
//参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

type database interface {
	connect() //只要实现了这个方法的变量都是database类型
	//execute(string,int)int
}

type mysql struct{}

type oracel struct{}

func (m mysql) connect() {
	fmt.Println("mysql连接方式")
}

func (m mysql) delete() {
	fmt.Println("mysql删除")
}

func (o oracel) connect() {
	fmt.Println("oracel连接方式")
}

// func connect(d database) {
// 	d.connect()
// }

func main() {
	//接口类型变量能够存储所有实现了该接口的实例。
	var db database
	db = mysql{}
	fmt.Printf("%T\n", db) //main.mysql
	fmt.Printf("%T\n", db) //main.mysql
	db.connect()

	// m := mysql{}
	// var db database
	// fmt.Printf("%T\n", db) //<nil>
	// db = m
	// fmt.Printf("%T\n", db) //main.mysql
	// db.connect()
}
