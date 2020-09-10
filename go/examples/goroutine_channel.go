package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建带缓存的通道
	buf := 100
	ch := make(chan string, buf)

	fmt.Println(ch) // 0xc00004a060

	// 使用for 或者 for-range 遍历通道,会自动检测通道是否关闭
	// for v := range ch {
	// 	fmt.Println(v)
	// }

	// 检测通道是否关闭
	// for {
	// 	v, open := <-ch
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(v)
	// }

	// 通过一个通道让主程序等待,直到协程完成: 信号量模式
	ch1 := make(chan int)
	go func() {
		const v int = 1
		fmt.Println(v) // 1
		// doSomething
		time.Sleep(2e9)
		ch1 <- v
	}()
	doSomethingElseForAWhile()
	// <-ch1

	// 终止一个协程
	// runtime.Goexit()

	// 简单的超时模板
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()
	select {
	case v := <-ch1:
		fmt.Println("ch1", v)
	case v := <-timeout:
		fmt.Println("timeout", v)
	}
}

func doSomethingElseForAWhile(params ...interface{}) {}

// 通道的工厂模式: 启动一个匿名函数作为协程生产通道
func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 使用输入通道和输出通道代替锁
type Task int

func Worker(in, out chan *Task) {
	for {
		t := <-in
		doSomethingElseForAWhile(t)
		out <- t
	}
}
