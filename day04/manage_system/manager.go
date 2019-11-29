package main

import "fmt"

type student struct {
	id   int
	name string
}

type manager struct {
	members []student
}

func (m manager) fetch() {
	for _, v := range m.members {
		fmt.Printf("学号:%d,姓名:%s\n", v.id, v.name)
	}
}

func (m *manager) add() {
	var (
		id   int
		name string
	)
	fmt.Println("请输入学号:")
	fmt.Scanln(&id)
	fmt.Println("请输入姓名:")
	fmt.Scanln(&name)
	s := student{
		id:   id,
		name: name,
	}
	m.members = append(m.members, s)
}

func (m manager) del() {

}
