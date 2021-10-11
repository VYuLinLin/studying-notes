package main

import (
	"fmt"
	"os"

	"./parse"
)

func badCall() {
	panic("bad end")
}
func test() {
	defer func() {
		fmt.Println("defer recover")
		if e := recover(); e != nil {
			fmt.Printf("panicing %s\r\n", e)
		}
	}()
	badCall()
	// 因为上面有调用recover，所以下面代码不会执行
	fmt.Printf("after bad call\r\n")
}
func main() {
	fmt.Println(os.Getenv("USER"))
	fmt.Println("starting the program")
	// panic("A severe error occurred: stopping the program!")
	test()
	fmt.Println("endding the program")

	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6 6.5",
		"2 + 3 = 5",
		"23swf",
		"",
	}

	for _, ex := range examples {
		fmt.Println("parsing ", ex)
		nums, err := parse.Parse(ex)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(nums)
	}
}
