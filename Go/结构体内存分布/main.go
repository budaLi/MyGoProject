package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

type React struct {
	left, right Point
}

func main() {
	r1 := React{Point{1, 2}, Point{3, 4}}
	fmt.Println(&r1.left.x, &r1.left.y, &r1.right.x, &r1.right.y)
	//可以看到结构体中的变量 在内存中的地址是连续的
	//0xc000060140 0xc000060148 0xc000060150 0xc000060158

	//在这里可以顺便学习下16进制的加减法
	//可以看到每个变量在内存中的地址差8个字节 内存中表示地址为16进制
	//而8用16进制表示仍然为8   1 2 3 4 5 6 7  8 9 A B C D E F
	//   60140
	//  +    8
	//  _______
	//   60148

	//  60148
	// +    8
	//————————
	//  60150
	
	
	//满16进1即可