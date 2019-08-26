package main

import (
	"fmt"
	"sync"
	"time"
)

//编写一个函数 实现1-200各个树阶乘 并将其放入到map中
//使用多个协程进行
//map应该定义为全局的

var myMap = make(map[int]int, 10)
var lock sync.Mutex

//sync 是同步 synchornized
//Mutex是互斥

//返回一个数的阶乘结果并将其放入到map中
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()

}
func main() {
	for i := 1; i <= 20; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 5)
	for i, v := range myMap {
		fmt.Printf("Map[%d]:%d", i, v)
		fmt.Println()
	}
}
