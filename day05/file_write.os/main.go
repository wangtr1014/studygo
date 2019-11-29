package main

import (
	"fmt"
	"os"
)

func main() {
	//os.O_TRUNC用到了清空 所以每次打开文件都先清空再写
	//追加用os.O_APPEND
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello 沙河"
	file.Write([]byte(str))       //写入字节切片数据
	file.WriteString("hello 小王子") //直接写入字符串数据
}
