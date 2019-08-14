package main

import (
	"fmt"
)

func BubbleSort(arr *[5]int) {
	fmt.Println("未排序数组为", (*arr))
	for i := 1; i <= len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}
func main() {
	arr := [5]int{1, 9, 45, 23, 3}
	BubbleSort(&arr)
	fmt.Println(arr)
}
