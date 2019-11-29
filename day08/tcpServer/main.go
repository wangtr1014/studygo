package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handle(conn net.Conn) {
	defer conn.Close() // 关闭连接
	addr := conn.RemoteAddr()
	fmt.Println(addr)
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("创建服务端失败:", err)
		os.Exit(1)
	}
	defer l.Close()
	//持续等待连接
	for {
		conn, err := l.Accept() //等待连接
		if err != nil {
			fmt.Println("连接失败:", err)
			continue
		}
		go handle(conn)
	}

}
