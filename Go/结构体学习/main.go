package main

import (
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Score [5]int
	Per   *int
	Silce []int
	Map1  map[string]string
}

func main() {
	var person Person
	//当指针，切片，map为结构体的属性时 默认为nil 常规类型如int默认为0 string 为“”
	fmt.Println(person)
	person.Name = "力不大"
	person.Age = 17
	person.Score = [5]int{1, 2, 3, 4, 5}
	person.Silce = make([]int, 2)
	person.Silce[0] = 1
	person.Map1 = make(map[string]string)
	person.Map1["宝贝"] = "嘟嘟"
	fmt.Println(person)

	//结构体对象传值默认为值传递 也就是说 当结构体对象赋值时  相当于重新创建了一个对象 对该对象的改变不会对原对象有影响
	//当然 结构体对象复制时也可以使用引用传递 需要使用指针 如person1:=&person

	//结构体声明的几种方式 var person Person;  person2:=Person{}
	person1 := person
	fmt.Println(person1)
	person1.Name = "嘟嘟"
	fmt.Println(person1)

	//也可以这样声明一个结构体对象 这时p3为一个指针 那么赋值时需要注意
	p3 := new(Person)
	// var p3 *Person = new(Person)  和上面那种方式一样 更容易理解
	(*p3).Name = "李东八不懂得"
	fmt.Println((*p3))

	//这种写法也可以
	var p4 *Person = &Person{}
	fmt.Println(p4)
}
