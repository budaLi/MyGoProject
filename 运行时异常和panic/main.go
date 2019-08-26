package main

import (
	"fmt"
)

//当发生像数组下标越界或者类型断言失败这样的错误时 go运行时会触发运行Panic
//伴随着程序的奔溃跑出一个runtime.Error接口类型的值
//panic可以直接从代码初始化 当运行条件很严苛且不可恢复，程序不能继续运行时
//可以调用panic函数产生一个终止程序的运行时错误 panic接受一个任意类型的参数
//通常是字符串 在程序死亡时被打印出来

func main() {
	fmt.Println("程序开始执行")
	panic("ERROR")
	fmt.Println("程序结束")
}
