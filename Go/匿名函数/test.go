package main

import (
	"fmt"
	"math"
)

func main() {
//	//1.定义时调用匿名函数
//	func(data int) {
//		fmt.Println("定义时调用", data)
//	}(100)

//	//2.定义时调用有参匿名函数
//	res := func(data float64) float64 {
//		return math.Sqrt(data)
//	}(250)
//	fmt.Println(res)

//	//3.将匿名函数赋值给变量 需要时调用
//	f := func(data string) string {
//		return data
//	}
//	fmt.Println(f("将匿名函数赋值给变量"))

	//4.匿名函数作为回调函数
	
	func FilterSlices(a []float64 , f myFunc) (tem []string){

		for _,value := range a {
			tem = append(tem,f(value))	
		}
		return tem
		}
		
		type myFunc func(float64) string
		var arr= []float64{1,9,16,25}
		var result= FilterSlices(arr,func(val float64) (string){
		val = math.Sqrt(val)
		return fmt.Sprintf("%.0f",val)
		})
		fmt.Println("匿名函数作为回调函数",result)
		
		//遍历切片 对其中每个元素进行运算处理


}
