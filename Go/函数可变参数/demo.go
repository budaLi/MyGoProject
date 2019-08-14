package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 4, 5, 6}
	fmt.Println(test(1, 2, 3))
	fmt.Println(test(arr...))
}

func test(arr ...int) (sum int) {
	for _, value := range arr {
		sum += value
	}
	return sum
}
