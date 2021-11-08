package main

import (
	"fmt"
	"testing"
)

// go函数的性能基准测试
func main() {
	ch := make(chan int) // Allocate a channel.
	// Start something in a goroutine; when it completes, signal on the channel.
	go func() {
		// doSomething
		ch <- 1 // Send a signal; value does not matter.
	}()
	// doSomethingElseForAWhile()
	fmt.Println(" sync", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("buffered", testing.Benchmark(BenchmarkChannelBuffered).String())
	<-ch

}

func BenchmarkChannelSync(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}

func BenchmarkChannelBuffered(b *testing.B) {
	ch := make(chan int, 128)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}
