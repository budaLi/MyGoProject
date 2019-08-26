package main

import (
	"errors"
	"fmt"
	"math"
)

//go中有个预先定义的error接口类型
// type error interface{
//		Error() string
//}  errors包中有一个errorString 结构体实现了该接口
var err error = errors.New("这是一个报错信息")

func main() {
	fmt.Println(err)

	//平方根参数测试函数
	res, errs := Sqrt(-1)
	if err != nil {
		fmt.Println(errs)
	} else {
		fmt.Println(res)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("被开根数不能小于0")
	} else {
		return math.Sqrt(f), nil
	}
}
