package main

import (
	"fmt"
	"strings"
)

func main() {
	//双引号包裹的是字符串
	//单引号包裹的是字符
	//字节 Byte 1 Byte=8 Bit(8个二进制位)
	//一个字符'A'=一个字节
	//一个utf8编码的汉字'你'=一般占三个字节
	path := "/Users/wangtr/go/src/github.com/wangtr1014/studygo/day01/bool"
	fmt.Println(path)
	// \r回车符（返回行首）\n换行符（直接跳到下一行的同列位置）\t制表符
	// \'单引号 \"双引号 \\反斜杠
	path2 := "\"/Users/wangtr/go/src/github.com/wangtr1014/studygo/day01/bool\""
	fmt.Println(path2)
	//多行字符串用反引号（原样输出 无需转义）
	str := `
		123
			456
	`
	fmt.Println(str)
	//字符串拼接
	s1 := "ab"
	s2 := "cd"
	fmt.Println(s1 + s2)
	s3 := s1 + s2
	fmt.Println(s3)
	//直接打印
	fmt.Printf("%s%s\n", "hello", "world")
	//这是一个赋值函数
	s4 := fmt.Sprintf("%s%s", "hello", "world")
	fmt.Println(s4)

	//字符串操作
	str = "这|是|一|句|话"
	//长度
	fmt.Println(len(str))
	//分割
	ret := strings.Split(str, "|")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(str, "是"))
	//前缀
	fmt.Println(strings.HasPrefix(str, "这|是"))
	//后缀
	fmt.Println(strings.HasSuffix(str, "句|话"))
	//子串出现的位置
	fmt.Println(strings.Index(str, "|"))
	fmt.Println(strings.LastIndex(str, "|"))
	//拼接
	fmt.Println(strings.Join(ret, "+"))
}
