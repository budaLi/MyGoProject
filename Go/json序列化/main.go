package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Hobby string
}

func main() {
	person := Person{"libuda", 12, "吃饭"}
	fmt.Println(person)
	jsonstr, ok := json.Marshal(person)
	if ok != nil {
		fmt.Println("序列化失败")
	}
	fmt.Println(string(jsonstr))
}
