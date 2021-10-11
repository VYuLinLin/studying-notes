package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 当前的内存状态，单位KB
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc / 1024)
}
