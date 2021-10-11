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

	fmt.Printf("arr = %v, slice1 = %v \n", arr, slice1)

	// 使用new初始化slice不会分配内存地址，所以不能通过索引赋值，应该使用make初始化slice
	slice2 := new([]int)
	fmt.Printf("new created slice %v \n", slice2)
	fmt.Printf("array created slice %v \n", slice1)
}
