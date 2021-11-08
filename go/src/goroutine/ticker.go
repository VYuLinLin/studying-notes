package main

import (
	"fmt"
	"time"
)

// 限制处理频率
func main() {
	var dur time.Duration = 1e9
	chRate := time.Tick(dur) // 每秒滴答一次
	for i := range pump() {
		t := <-chRate                  // 此行会阻塞后面的代码，循环方法的处理频率为每秒执行一次
		fmt.Println("a tick...", i, t) // a tick... 0 2021-10-30 02:20:37.9484483 +0800 CST m=+1.003293301
		if i >= 100 {
			break
		}
	}
}
func pump() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}
