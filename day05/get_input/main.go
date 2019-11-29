package main

import (
	"bufio"
	"fmt"
	"os"
)

func useScan() {
	//加入输入的是a b c 只能拿到a
	//疑问Scanln遇到空格、回车都会停止接收
	fmt.Println("请输入内容：")
	var s string
	fmt.Scanln(&s)
	fmt.Printf("你输入的是%s\n", s)
}

func useBufio() {
	fmt.Println("请输入内容：")
	//从标准输入中读取内容可以读到完整内容
	reader := bufio.NewReader(os.Stdin)
	//按换行符结束赋值
	input, _ := reader.ReadString('\n')
	fmt.Printf("你输入的是%s\n", input)
}

func main() {
	//useScan()
	useBufio()
}
