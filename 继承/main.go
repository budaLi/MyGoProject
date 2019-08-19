package main

import (
	"fmt"
)

//继承可以解决代码复用 让我们的编程更接近人类思维
//当多个结构体中存在相同的属性和方法时 我么可以抽象出一个结构体 在其中声明相同的属性和方法
//在 go中 如果一个结构体嵌套加入了另一个结构体 那么它可以直接使用该结构体的字段和方法 从而实现了继承
//在这里要注意的是  嵌套其他结构体的结构体 可以使用该结构体的所有字段和方法 注意是所有 包括小写字母开头的属性、
//如果一个结构体中嵌套多个结构体 且这多个结构体拥有相同的属性和方法、该结构体没有时  必须指定该结构体 名字
//如果一个结构体中嵌套了一个有名结构体 那么在该结构体访问该有名结构体时 必须带上该结构体的名字 否则会编译出错
type Person struct {
	Name string
	age  int
}

type Student struct {
	work string
	Name string
	Person
}

type XiaoStudent struct {
	Person
	Student
}

func main() {
	stu := Student{}

	//通过以下两种方法都可以设置
	stu.Name = "buda"
	stu.Person.Name = "baba"
	fmt.Println(stu.Name)

	stu.age = 11
	fmt.Println(stu.age)

	xiaostu := XiaoStudent{}
	fmt.Println(xiaostu.age)
	xiaostu.work = "吃饭"
	fmt.Println(xiaostu.work)

}
