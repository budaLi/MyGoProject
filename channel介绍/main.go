package main

import (
	"fmt"
)

//channel本质上是一种数据结构 队列
//数据是先进先出
//线程安全 多groutine访问时 不需要枷锁 就是说channel本身就是线程安全的
//channel是有类型的 一个string类型的channel只能存放string类型

func main() {
	//声明一个管道
	var mychannel chan int
	//make空间
	mychannel = make(chan int, 3)

	fmt.Println(mychannel)

	//向管道写入数据
	//注意 当我们向管道写入数据时  不能超过其容量
	mychannel <- 11
	num := 2
	mychannel <- num

	//查看管道的容量和长度
	fmt.Println(len(mychannel))
	fmt.Println(cap(mychannel))

	//从管道中读取数据
	num2 := <-mychannel
	num3 := <-mychannel
	//	num4 := <-mychannel
	fmt.Println(num2)
	fmt.Println(num3)
	//	fmt.Println(num4)
	fmt.Println(len(mychannel))

	//在没有使用协程的情况下 如果我们的管道数据以及全部取出 再取的话会报错 deadlock
}
