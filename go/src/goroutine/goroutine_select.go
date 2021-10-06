package main

import (
	"fmt"
	"time"
)

func main() {
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// go pump1(ch1)
	// go pump2(ch2)
	// go suck(ch1, ch2)

	time.Sleep(1e9)

	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}

	// fmt.Println("Tick", time.Tick(1e9))    // Tick 0xc000128000
	// fmt.Println("Aflter", time.After(1e9)) // Tick 0xc000128000
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Println("通道1： ", v)
		case v := <-ch2:
			fmt.Println("通道2： ", v)
		}
	}
}
