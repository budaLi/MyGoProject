package main

import (
	"fmt"
)

//一个接口可以继承多个接口 但是当自定义类型要实现该接口时 也必须要实现该接口继承的其他接口
//接口类型默认是一个指针 如果对其没有初始化就使用 那么会输出nil
//空接口没有任何方法 所有类型都实现了空接口 即我们将其指向任何一个变量

//接口和继承解决的问题不同
//继承的价值在于 提高代码的复用性和可维护性
//接口的价值在于 设计 定制好规范 让其他自定义类型去实现这些方法
//接口比继承更加灵活 继承是满足 is a 的关系 而接口只需要满足like a 的关系
//接口在一定程度上实现了代码解耦
type A interface {
	testA()
}

type B interface {
	testB()
}

type C interface {
	testC()
	A
	B
}

type person struct {
}

func (p person) testA() {
	fmt.Println("a")
}

func (p person) testB() {
	fmt.Println("b")
}

func (p person) testC() {
	fmt.Println("c")
}

func main() {
	var p person
	var c C = p
	c.testA()
	c.testB()
	c.testC()

}
