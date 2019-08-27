package main

import (
	"fmt"
)

//使用内置函数close可以关闭channel 当channel关闭后
//就不能向channel写数据 但是仍然可以读数据

func main() {
	mychan := make(chan int, 10)
	mychan <- 10
	mychan <- 200
	//	close(mychan)
	intnum := <-mychan
	fmt.Println(intnum)
	mychan <- 10

	mychan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		mychan2 <- i
	}

	close(mychan2)
	//在遍历时  如果channel没有关闭 就会出现deadlock的错误
	//如果channel已经关闭 就会正常遍历数据 遍历完后 就会正常退出遍历
	//遍历管道不能使用普通的for循环
	for v := range mychan2 {
		fmt.Println(v)
	}
}
