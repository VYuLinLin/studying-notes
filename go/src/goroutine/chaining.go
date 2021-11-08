package main

import (
	"flag"
	"math/big"

	"fmt"
)

var ngoroutine = flag.Int("n", 100000, "how many goroutines")

func f(left, right chan int) { left <- 1 + <-right }

// 链式操作
// go run chaining.go -n=70000
func main() {

	flag.Parse()

	leftmost := make(chan int)

	var left, right chan int = nil, leftmost

	for i := 0; i < *ngoroutine; i++ {

		left, right = right, make(chan int) // 链式操作

		go f(left, right)

	}

	right <- 0

	// start the chaining 开始链接

	x := <-leftmost // wait for completion 等待完成

	fmt.Println(x) // 100000

	// 结果： 100000 ， 大约 1,5s （我实际测试只用了不到200ms）

	var r, a, b, c big.Int
	a = *big.NewInt(7)
	b = *big.NewInt(42)
	c = *big.NewInt(24)

	// r = a * (b - c)
	r.Mul(&a, r.Sub(&b, &c)) // 链式操作

	fmt.Println(r.String(), b) // 126 {false [42]}
}
