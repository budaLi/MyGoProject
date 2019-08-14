package main

import (
	"fmt"
)

func main() {
	var cat Cat
	//当结构体声明时  其已经分配内存空间并复制 string类型默认为“” int 类型默认为0S
	fmt.Println(cat)
	cat.Name = "xiaobao"
	cat.Age = 14
	cat.Color = "red"
	fmt.Println("名字", cat.Name)
	fmt.Println("年龄", cat.Age)
	fmt.Println("颜色", cat.Color)
}

type Cat struct {
	Name  string
	Age   int
	Color string
}
