package main

import (
	"fmt"
)

type A struct {
	Num int
}

func (a A) test() {
	a.Num = 100
	fmt.Println(a.Num)
}

func main() {
	var a A
	a.test()
}
