package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "默认为空")
	flag.StringVar(&pwd, "pwd", "", "默认为空")
	flag.StringVar(&host, "h", "localhost", "默认为localhost")
	flag.IntVar(&port, "port", 3306, "默认为3306")

	flag.Parse()

	//当命令中 输入go run main.go -u root -pwd root -h loc -port 22
	//打印出root root loc 22
	fmt.Println(user, pwd, host, port)
}
