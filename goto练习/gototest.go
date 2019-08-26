package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("打印1-100以内的所有素数")
	printSuShu()
}

func printSuShu() {
	i := 1
LOOP:
	for i <= 100 {
		for j := 2; j < int(math.Sqrt(float64(i)))+1; j++ {
			if i%j == 0 {
				goto LOOP
			}
		}
		//		fmt.Println(i)

	}
}
