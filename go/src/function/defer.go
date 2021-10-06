package main

import (
	"fmt"
	"log"
	"runtime"
)

func where() {
	// 使用闭包打印，进行调试
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d", file, line)
}

var i int
var a = [...]string{"5646456"}

func main() {
	fmt.Printf("main start %d - %T - %v\n", i, a, a)
	func1()
	i += 1
	fmt.Printf("main end %d \n", i)
	// log.SetFlags(log.Llongfile)
	// log.Println(time.Now())

}

func func1() (i int) {
	i += 1
	fmt.Printf("start %d \n", i)
	defer func2()
	i += 1
	// where()
	fmt.Printf("end %d \n", i)
	return i
}

func func2() {
	i += 1
	fmt.Printf("I'm func2 %d \n", i)
	// where()
}
