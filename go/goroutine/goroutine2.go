package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	defer close(ch)
	go sendData(ch)
	go getData(ch)
	// sendData(ch) // fatal error: all goroutines are asleep - deadlock!
	// getData(ch)
	fmt.Println("main1")
	time.Sleep(1e9)
	fmt.Println("main2")
	// fmt.Println("\nmain: ", <-ch)
	for {
		s, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("\nmain: ", s)
		if s == "end channel" {
			return
		}
	}
}

func sendData(ch chan string) {
	fmt.Println("sendData1")

	ch <- "hello"
	ch <- " "
	ch <- "world"
	ch <- "! "
	ch <- "Manila"

	fmt.Println("\nsendData2")
}

func getData(ch chan string) {
	fmt.Println("getData1")

	// time.Sleep(2e9)
	for {
		input, ok := <-ch
		if ok {
			fmt.Printf(input)
		}
		if input == "Manila" {
			ch <- "end channel"
		}
	}

	fmt.Println("\ngetData2") // 因为上面的循环是无限循环，所以此行代码永远不会执行
}
