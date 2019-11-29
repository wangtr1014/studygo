package main

import "fmt"

func main() {
	//map是引用类型 map[KeyType]ValueType
	var m map[string]int           //还没有初始化，没有在内存中开辟空间 不能进行赋值
	fmt.Println(m == nil)          //true
	m1 := make(map[string]int, 10) //尽量预估容量，避免动态扩容
	fmt.Println(m1)
	m1["key"] = 10
	m1["value"] = 100
	fmt.Println(m1)
	fmt.Println(m1["value"])
	//fmt.Println(m1["aaa"]) 不存在的Key直接拿到对应数据类型的默认值
	//判断map中键是否存在
	v, ok := m1["test"]
	if !ok {
		fmt.Println("不存在的Key")
	} else {
		fmt.Println(v)
	}
	//遍历 遍历map时的元素顺序与添加键值对的顺序无关。
	//如果要按key排序，先取出所有key，排序后再循环取值
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//只取key
	for k := range m1 {
		fmt.Println(k)
	}
	//直接初始化
	userInfo := map[string]string{
		"username": "wtr",
		"password": "wtr",
	}
	fmt.Println(userInfo)
	//删除key
	delete(userInfo, "username")
	//删除不存在的key没有任何影响
}
