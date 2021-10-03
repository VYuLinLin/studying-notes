package main

import "fmt"

var i int

func main() {
	fmt.Printf("main start %d \n", i)
	func1()
	fmt.Printf("main end %d \n", i)
}

func func1() (i int) {
	i += 1
	fmt.Printf("start %d \n", i)
	defer func2()
	i += 1
	fmt.Printf("end %d \n", i)
	return
}

func func2() {
	i += 1
	fmt.Printf("I'm func2 %d \n", i)
}
