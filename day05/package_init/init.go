package initpackage

import "fmt"

func init() {
	fmt.Println("init包的init函数")
}

func Test() {
	fmt.Println("这里是init包的Test函数")
}
