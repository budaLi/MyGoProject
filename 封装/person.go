package main

import (
	"fmt"
)

type Person struct {
	Name string
	age  int //结构体中变量小写表示其他包中无法访问
	sal  float64
}

//工厂模式的函数  相当于构造函数 传入一个name生成一个person的对象
func NewPerson(name string) *Person {
	return &Person{
		Name: name,
	}
}

//get方法使能访问age
func (p *Person) GetAge() int {
	return p.age
}

//set方法使能够设置age
func (p *Person) SetAge(age int) {
	if age > 0 && age < 100 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}

func main() {
	p := Person{}
	age := p.GetAge()
	fmt.Println(age)
	p.SetAge(11)
	fmt.Println(p.GetAge())
}
