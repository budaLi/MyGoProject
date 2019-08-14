package main

import (
	"fmt"
)

func main() {
	//声明一个map 并未分配空间
	var a map[string]string
	//通过make分配内存
	a = make(map[string]string, 10)
	a["爸爸"] = "力不大"
	fmt.Println(a)
}
