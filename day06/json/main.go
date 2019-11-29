package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var p person
	str := `{"name":"wtr","age":25}`
	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(p)
	s, _ := json.Marshal(p)
	fmt.Println(string(s))
}
