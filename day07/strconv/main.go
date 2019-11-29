package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "100"
	//string 转int
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}
	i2 := 200
	//int转string
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"

	//字符串转换为数字
	// 10进制 32位
	i3, _ := strconv.ParseInt(s1, 10, 32)
	fmt.Printf("type:%T value:%#v\n", i3, i3) //type:int value:100

	//从字符串解析出bool值
	s3 := "true"
	b, _ := strconv.ParseBool(s3)
	fmt.Printf("type:%T value:%#v\n", b, b) //type:bool value:true

	//数字转换成字符串
	i := 500
	s4 := fmt.Sprintf("%d", i)
	fmt.Printf("type:%T value:%#v\n", s4, s4) //type:string value:"500"

	//ParseBool()、ParseFloat()、ParseInt()、ParseUint()
}
