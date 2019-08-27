package main

import (
	"fmt"
)

//使用goroutine和channel协同完成工作 具体要求
//开启一个writeData协程 向管道intchan中写入数据 50个整数
//开启一个readData协程 读取intchan中的内容
//注意 两个协程操作的是同一个管道
//主线程需要等待两个协程完成工作后才能退出

//如果编译器发现一个管道只有写没有读 当到该容量极限时会阻塞
//而当读管道和写管道的评率不一致时 无所谓 编译器会做出优化
func writeData(inchan chan int) {
	for i := 0; i < 50; i++ {
		inchan <- i
		fmt.Printf("writeData %v\n", i)
	}
	//写完后关闭管道
	close(inchan)
}

func readData(inchan chan int, exitchan chan bool) {
	for {
		v, ok := <-inchan
		if !ok {
			break
		}
		fmt.Printf("readData读取数据%v\n", v)
	}
	exitchan <- false
	close(exitchan)
}

func main() {
	intchan := make(chan int, 50)
	exitchan := make(chan bool, 1)

	go writeData(intchan)
	go readData(intchan, exitchan)

	for {
		_, ok := <-exitchan
		if !ok {
			break
		}
	}

}
