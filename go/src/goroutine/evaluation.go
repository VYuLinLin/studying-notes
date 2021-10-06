package main

import (
	"fmt"
	"time"
)

var resume chan int

func integers() chan int {
	yield := make(chan int)
	// yield := make(chan int, 34)
	count := 0
	go func() {
		for {
			yield <- count
			count++
			fmt.Println("inter", count)
		}
	}()
	return yield
}

func generateInteger() int {
	return <-resume
}

func main() {
	resume = integers()
	fmt.Println(generateInteger()) // 0
	fmt.Println(generateInteger()) // 1
	fmt.Println(generateInteger()) // 2
	fmt.Println(len(resume))       // 0  因为通道定义时没有设置缓冲数量，所以是0
	time.Sleep(1e9)
}
