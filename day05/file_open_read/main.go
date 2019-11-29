package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	//关闭文件
	defer fileObj.Close()
	//Read func (f *File) Read(b []byte) (n int, err error)
	//它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回0和io.EOF
	// var tmp [128]byte
	// fileObj.Read(tmp[:])
	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := fileObj.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}
