package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.1415926
	mapAssigned["two"] = 3

	fmt.Println("mapLit[\"one\"]", mapLit["one"])           // 1
	fmt.Println("mapCreated[\"key2\"]", mapCreated["key2"]) // 3.1415926
	fmt.Println("mapAssigned[\"two\"]", mapAssigned["two"]) // 3
	fmt.Println("mapLit[\"two\"]", mapLit["two"])           // 3
	fmt.Println("mapLit[\"ten\"]", mapLit["ten"])           // 0
	fmt.Println("mapLit", mapLit)                           // mapLit map[one:1 two:3]

	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}

	fmt.Println(mf)                   // map[1:0x49e970 2:0x49e980 5:0x49e990]
	fmt.Println(mf[1]())              // 10
	fmt.Println(reflect.TypeOf(0))    // int
	fmt.Println(reflect.TypeOf(0.00)) // float64
	fmt.Println(0 == 0.00)            // true
	fmt.Println(0 == '0')             // false

	mapSlice()

	mapSort()
}

func mapSlice() {
	items := make([]map[int]int, 5)

	fmt.Println(items) // [map[] map[] map[] map[] map[]]
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	// items =  [map[1:2] map[1:2] map[1:2] map[1:2] map[1:2]]
	fmt.Println("items = ", items)
}

func mapSort() {
	var (
		barVal = map[string]int{
			"monday":    1,
			"tuesday":   2,
			"wednesday": 3,
			"thursday":  4,
			"friday":    5,
			"saturday":  6,
			"sunday":    7,
		}
	)
	fmt.Println("默认排序：")
	for k, v := range barVal {
		fmt.Printf("key: %s， value:%d\n", k, v)
	}

	keys := make([]string, len(barVal))
	i := 0
	for k := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("根据key排序：")
	for _, k := range keys {
		fmt.Printf("key: %s， value:%d\n", k, barVal[k])
	}

	values := make([]int, len(barVal))
	j := 0
	for _, v := range barVal {
		values[j] = v
		j++
	}
	sort.Ints(values)
	fmt.Println("根据值排序：")
	for _, v := range values {
		for k, val := range barVal {
			if v == val {
				fmt.Printf("key: %s， value:%d\n", k, val)
			}
		}
	}
}
