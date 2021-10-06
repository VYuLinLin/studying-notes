package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [19]int{1, 2, 3, 4, 5, 6, 7, 743, 3, 54, 5, 4, 6, 43, 6, 3, 57, 34, 573}

	// sort.Ints(arr)
	fmt.Println(arr)
	fmt.Println(sort.SearchInts(arr[:], 7))
	var i int = 8
	a := append(arr[:7], append([]int{i}, arr[7:]...)...)
	fmt.Println(a)
}
