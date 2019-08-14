package main

import "fmt"

const (
	a = iota
	b = 8
	c
)

//通过Iota可实现存储单元的常量枚举 iota从0开始
type ByteSize float64

const (
	_           = iota //通过_可忽略0值
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {

	//新的常量b声明后 iota将不再往下赋值 后面的常量如果
	//没有赋值 将继承上一个常量值
	fmt.Println(a, b, c)
	fmt.Println(KB, MB, GB, TB)
}
