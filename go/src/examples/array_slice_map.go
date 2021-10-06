package main

import "fmt"

func main() {
	// 创建
	arr := new([3]string)
	slice := make([]int16, 3)
	fmt.Println(arr)   // &[  ]
	fmt.Println(slice) // [0 0 0]

	// 初始化数组、切片
	arr1 := [...]int32{1, 2, 3, 4}
	var slice1 []int32 = arr1[0:]
	fmt.Println(arr1)   // [1 2 3 4]
	fmt.Println(slice1) // [1 2 3 4]

	// 截断数组或者切片的最后一个元素
	fmt.Println(arr1[:len(arr1)-1]) // [1 2 3]

	// for 或者 for-range 遍历数组、切片
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	// 创建映射
	// map1 := make(map[keytype]valuetype)

	// 初始化
	map1 := map[string]int{"one": 1, "two": 2}

	// 遍历
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	// 检测map中是否包含key1
	if v, isPresent := map1["one"]; isPresent {
		fmt.Println(v) // 1
	}

	// 删除map中的键
	delete(map1, "one")
	fmt.Println(map1) // map[two:2]
}
