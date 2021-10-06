package main

import "fmt"

func main() {
	s := make([]int, 5)
	fmt.Printf("len = %d, cap = %d \n", len(s), cap(s)) // len = 5, cap = 5

	s = s[2:4]
	fmt.Printf("len = %d, cap = %d \n", len(s), cap(s)) // len = 2, cap = 3

	// 切片的值修改会改变相关数组的值
	var arr [10]int
	slice1 := arr[:]
	slice1[2] = 222

	fmt.Printf("arr = %v, slice1 = %v", arr, slice1)
}
