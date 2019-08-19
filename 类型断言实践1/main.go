package main

//类型断言的实践
//给phone结构体 增加一个新的方法call()  当usb接口接受的是Phone变量时 还需要调用call方法
import (
	"fmt"
)

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

func (p Phone) Call() {
	fmt.Println("手机在叫")
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
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	phone := Phone{}
	camara := Camara{}
	//声明一个Usb数组 可以存放phone变量 和camaea变量

	var Usbarr [2]Usb
	Usbarr[0] = phone
	Usbarr[1] = camara

	//声明一个computer对象
	var comp Computer
	for _, v := range Usbarr {
		comp.Working(v)
		fmt.Println()
	}
}
