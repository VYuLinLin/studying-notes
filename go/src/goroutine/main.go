package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start main()")
	// 协程以关键字go开头的调用实现
	go longWait()
	go shortWait()
	// longWait()
	// shortWait()
	fmt.Println("About to sleep in main()")
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")

	// 习惯用法： 通道工厂模式
	stream := pump()
	go suck(stream)
	// 或
	// go suck(pump())
	time.Sleep(1e9)
}
func longWait() {
	fmt.Println("Beginning longWait")
	time.Sleep(5 * 1e9)
	fmt.Println("End of longWait")
}
func shortWait() {
	fmt.Println("Beginning shortWait")
	time.Sleep(2 * 1e9)
	fmt.Println("End of shortWait")
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
