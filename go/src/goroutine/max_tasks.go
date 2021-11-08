package main

import (
	"fmt"
	"time"
)

const (
	AvailableMemory = 10 << 20

	// 10 MB, 示例

	AverageMemoryPerRequest = 10 << 10

	// 10 KB

	MAXREQS = AvailableMemory / AverageMemoryPerRequest
	// 原文中说 MAXREQS 是 1000，实际计算是 1024 ，后面按照原文的 1000 来描述

)

var sem = make(chan int, 2)

type Request struct {
	a, b int

	replyc chan int
}

func process(r *Request) {

	// Do something 做任何事

	// 可能需要很长时间并使用大量内存或CPU
	time.Sleep(1e9)
	fmt.Println("process end")
}

func handle(r *Request) {

	process(r)

	// 信号完成：开始启用下一个请求

	// 将 sem 的缓冲区释放一个位置

	<-sem
}

func Server(queue chan *Request) {

	for {

		sem <- 1

		// 当通道已满（1000 个请求被激活）的时候将被阻塞

		// 所以停在这里等待，直到 sem 有容量（被释放），才能继续去处理请求

		// (doesn’t matter what we put in it)

		fmt.Println("Server")
		request := <-queue

		go handle(request)

	}

}

func main() {

	queue := make(chan *Request)

	go Server(queue)

	for i := 0; i < 10; i++ {
		queue <- &Request{i, i * 2, make(chan int)}
	}

	fmt.Println("main end")
}
