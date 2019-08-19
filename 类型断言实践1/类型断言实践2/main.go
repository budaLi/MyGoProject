package main

import (
	"fmt"
)

//写一个函数 循环判断传入参数的类型

func JudgeType(item ...interface{}) {
	for _, x := range item {
		switch x.(type) {
		case bool:
			fmt.Printf("%v为bool类型\n", x)
		case int:
			fmt.Printf("%v为int类型\n", x)
		case string:
			fmt.Printf("%v为string类型\n", x)
		default:
			fmt.Printf("%v为其他类型\n", x)
		}
	}
}

func main() {
	var n1 float64 = 11
	var n2 int = 1
	var n3 string = "力不大"
	var n4 bool = false

	JudgeType(n1, n2, n3, n4)
}
