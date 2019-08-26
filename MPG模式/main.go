package main

import (
	"fmt"
	"runtime"
)

//M:操作系统的主线程 是物理线程
//P:协程执行所需的上下文 可以理解为协程执行所需的环境
//G：协程

func main() {
	//得到当前系统的cpu数目
	num := runtime.NumCPU()
	fmt.Println(num)

	//可以自己设置使用多少个cpu
	runtime.GOMAXPROCS(2)
}
