package main

import (
	"fmt"
)

//如何讲一个接口变量赋值给一个类型变量 在这里就需要类型断言

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	a = Point{1, 2}
	var b Point
	//	b = a //在这里会报错 cannot use a (type interface {}) as type Point in assignment: need type assertion

	//由于接口是一般类型 不知道其具体类型 那么当它要转成具体类型 就要使用类型断言 在进行类型断言时 要确保接口原来指向的类型就是当前断言的类型

	b = a.(Point)
	fmt.Println(b)

	//带判断的类型断言 从下面代码可以看出 类型断言函数 返回两个值 第一个是类型断言后的对象 第二个为断言的结果 我们可以根据第二个参数进行判断
	if y, ok := a.(Point); ok {
		fmt.Println("success", y)
	} else {
		fmt.Println("error")
	}

}
