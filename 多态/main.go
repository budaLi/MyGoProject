package main

import (
	"fmt"
)

//在go中  多态是通过接口体现的
//为什么要有接口 那么在生活中 电脑上只有一个接口 但是同时可以接鼠标 U盘等多个东西
//这是因为生产厂商在定制usb接口时使用了相同的规范

//接口可以定义一组方法 但是这些不需要实现 并且接口不能包含任何变量 当某个结构体类型
//要实现该接口中的方法时 再根据具体情况将函数实现

//小结
//接口中的所有方法都没有方法体 即接口中的方法都是没有实现的方法 接口体现了程序设计中高内聚低耦合 和 多态的特点
// go中的接口不需要显示的实现 只需要一个结构体对象 实现了该接口的所有方法	那么这个对象就实现了这个接口 因此
// go中没有Implement这样的关键字
type Usb interface {
	Start() //声明了两个没有实现的方法
	Stop()
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Camara struct {
}

func (cam Camara) Start() {
	fmt.Println("相机开始工作")
}

func (cma Camara) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

func (com Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	phone := Phone{}
	//	canara := Camara{}
	computer := Computer{}
	computer.Working(phone)
}
