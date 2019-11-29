package main

import "fmt"

func main() {
	s := "hello世界"
	//len取得的是byte字节的长度
	l := len(s)
	fmt.Println(l)
	//for循环获取字符 byte
	for i := 0; i < l; i++ {
		fmt.Printf("%c\n", s[i])
	}
	fmt.Println("---------------------")
	//range获取rune rune类型代表一个UTF-8字符
	for i, r := range s {
		fmt.Printf("%d-%c\n", i, r)
	}
	//字符串底层是一个byte数组，所以可以和[]byte类型相互转换。
	//字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。
	//rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。

	//字符串修改
	//要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。
	//无论哪种转换，都会重新分配内存，并复制字节数组。
	s1 := "白萝卜"
	s2 := []rune(s1) //字符串强制转换为rune切片
	s2[0] = '绿'      //注意这里是字符类型
	fmt.Println(s2)
	fmt.Println(string(s2)) //rune切片强制转换为string
	s3 := "hello"
	s4 := []byte(s3) //字符串转byte切片
	s4[0] = 'w'
	fmt.Println(s3)
	fmt.Println(string(s4)) //byte切片强制转换为string

	//类型转换
	n := 10
	f := float32(n)
	fmt.Printf("%T\n", f)
}
