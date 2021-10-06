package main

import (
	"fmt"
	"time"
)

// 信号量模式
// 通常做法是在main函数的最后放置一个{}
// 或者使用通道让main程序等待协程完成
func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("111")
		time.Sleep(1e9)
		ch <- 1
	}()
	fmt.Println("222")
	<-ch
	fmt.Println("333")
}
