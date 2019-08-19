package main

import (
	"fmt"
)

//接口的几个应用场景
//1  程序开发中

//注意事项
//一个自定义类型只有实现了一个接口的所有方法 我们才说这个自定义类型实现了接口 此时才能将该自定义类型的实例对象赋给 接口类型 也就是说
//接口本身不能创建实例 但是可以指向一个实现了该接口所有方法的自定义的对象

type Inface interface {
	Say()
}

type Child struct {
}

func (child Child) Say() {
	fmt.Println("我在说")
}

//只要是自定义类型 就可以实现结构 而不仅仅是结构体类型
type strs string

func (st strs) Say() {
	fmt.Println("字符串再说")
}

func main() {

	child := Child{}
	var infa Inface = child
	infa.Say()

	var st strs = "buda"
	infa = st
	infa.Say()
}
