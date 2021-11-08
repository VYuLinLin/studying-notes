package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 典型的 - 生产者消费者模式
	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e9)

	// 超时模式
	tick := time.Tick(1e9)
	boom := time.After(5e9)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
			// default:
			// 	fmt.Println("    .")
			// 	time.Sleep(5e7)
		}
	}

	// fmt.Println("Tick", time.Tick(1e9))    // Tick 0xc000128000
	// fmt.Println("Aflter", time.After(1e9)) // Tick 0xc000128000
}
func pump1(ch chan int) {
	flag := false
	go func() {
		time.Sleep(1e9)
		flag = true
	}()
	for i := 0; ; i++ {
		if flag {
			close(ch)
			return
		}
		time.Sleep(5e8)
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	flag := false
	go func() {
		time.Sleep(1e9)
		flag = true
	}()
	for i := 0; ; i++ {
		if flag {
			close(ch)
			return
		}
		time.Sleep(5e8)
		ch <- i + 5
	}
}

func suck(ch1, ch2 <-chan int) {
	for {
		select {
		case v, ok := <-ch1:
			if !ok {
				continue
			}
			fmt.Printf("通道1： %d %v \n", v, ok)
		case v, ok := <-ch2:
			if !ok {
				continue
			}
			fmt.Printf("通道2： %d %v \n", v, ok)
		}
	}
}
