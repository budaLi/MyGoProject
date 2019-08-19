package main

import (
	"fmt"
)

//一个自定义类型可以实现多个接口
//接口中不能存在变量

type Aintefac interface {
	Say()
}

type Bintefac interface {
	Eat()
}

type perple struct {
}

func (p perple) Say() {
	fmt.Println("人在说")
}

func (p perple) Eat() {
	fmt.Println("人在吃")
}

func main() {
	var p perple
	var aint Aintefac = p
	var bint Bintefac = p

	aint.Say()
	bint.Eat()

}
