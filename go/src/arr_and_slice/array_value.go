package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr1 := new([5]int)
	var arr2 [5]int
	var a []int
	fmt.Printf("arr1 type: %T - %s \n", arr1, reflect.TypeOf(arr1)) // arr1 type: *[5]int - *[5]int 数组
	fmt.Printf("a type: %T - %s \n", a, reflect.TypeOf(a))          // a type: []int - []int 切片

	arr11 := *arr1
	arr11[0] = 111

	arr22 := arr2
	arr22[0] = 222

	arr222 := &arr2
	arr222[0] = 2222

	fmt.Printf("arr1 = %d \n", arr1)                                               // arr1 = &[0 0 0 0 0]
	fmt.Printf("arr2 = %d \n", arr2)                                               // arr2 = [2222 0 0 0 0]
	fmt.Printf("arr11 = %d \n", arr11)                                             // arr11 = [111 0 0 0 0]
	fmt.Printf("arr22 = %d \n", arr22)                                             // arr22 = [222 0 0 0 0]
	fmt.Printf("arr222 = %d - %T - %s \n", arr222, arr222, reflect.TypeOf(arr222)) // arr222 = &[2222 0 0 0 0] - *[5]int - *[5]int
}
