package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	//可以存放任意数据类型的channel
	allchan := make(chan interface{}, 10)
	allchan <- 10
	allchan <- "string"
	person := Person{"力不大", 11}
	allchan <- person

	<-allchan
	<-allchan

	newperson := <-allchan
	//在这里要注意的是 newperson为interface类型 需要使用类型断言才能访问其属性
	fmt.Println(newperson)
	p := newperson.(Person)
	fmt.Println(p.Name)
}
