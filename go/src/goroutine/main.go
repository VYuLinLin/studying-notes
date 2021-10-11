package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 5; i++ {
		go fmt.Println("i = ", i)
	}
	fmt.Println("start main()")
	var n int
	// 协程以关键字go开头的调用实现
	go longWait(n)
	go shortWait(n)
	n = 6666
	// longWait()
	// shortWait()
	fmt.Println("About to sleep in main()")
	time.Sleep(4 * 1e9)
	fmt.Println("At the end of main()")

	// 习惯用法： 通道工厂模式
	stream := pump()
	go suck(stream)
	go suck1(stream)
	// 或
	// go suck(pump())
	time.Sleep(1e9)
}
func longWait(n int) {
	fmt.Println("Beginning longWait", n) // Beginning shortWait 0
	time.Sleep(3 * 1e9)
	fmt.Println("End of longWait")
}
func shortWait(n int) {
	fmt.Println("Beginning shortWait", n) // Beginning longWait 0
	time.Sleep(2 * 1e9)
	fmt.Println("End of shortWait")
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch <-chan int) {
	go func() {
		for v := range ch {
			fmt.Println("suck", v)
		}
	}()
}
func suck1(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println("suck1", v)
		}
	}()
}
