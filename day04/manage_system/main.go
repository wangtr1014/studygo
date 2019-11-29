package main

import (
	"fmt"
)

func menu() {
	fmt.Println(`
		1.查看所有学生信息
		2.新增学生
		3.删除学生
		4.退出系统
		`)
}

func main() {
	var studentManager manager
	for {
		menu()
		fmt.Println("请输入你的选择，按回车键确认:")
		var choice int
		fmt.Scanln(&choice)
		// fmt.Printf("你输入的是:%d\n", choice)
		switch choice {
		case 1:
			//查看所有学生
			studentManager.fetch()
		case 2:
			//新增学生
			studentManager.add()
		case 3:
			//删除学生
		case 4:
			//退出
			goto breakTag
		default:
			fmt.Println("输入有误")
		}
	}
breakTag:
	fmt.Println("退出系统")
}
