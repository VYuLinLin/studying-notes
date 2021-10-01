package main

import (
	"fmt"
	"strings"
	"time"
)

var num int = 10
var numx2, numx3 int

const LIM = 41

var fibs [LIM]uint64

func main() {
	numx2, numx3 = getTo2(num)
	printValues()
	numx2, numx3 = getTo3(num)
	printValues()

	n := 0
	reply := &n
	multiply(10, 5, reply)
	println("n = ", n)          // n =  50
	println("reply = ", *reply) // reply =  50

	x := min(1, 2, 6, 3, 6)
	println("x = ", x) // x =  1
	slice := []int{7, 8, 5, 6, 2}
	x = min(slice...)
	println("x = ", x)         // x =  2
	println("min() = ", min()) // min() =  0

	typecheck(5, 6, 2.3, 0.6, -6, -0.41, "hello", true)

	deferFunc1()

	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("\nfibonacci(%d) is: %d", i, result)
	}

	end := time.Now()
	delte := end.Sub(start)
	println("\n共消耗时间：", delte)

	// 匿名函数
	fplus := func(x, y int) int { return x + y }
	println("\n匿名函数赋值给变量：", fplus(2, 3))
	println("匿名函数直接调用：", func(x, y int) int { return x + y }(3, 6))

	println("匿名函数defer", fv()) // 2

	println(add2()(3)) // 5

	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")

	println(addBmp("file"))  // file.bmp
	println(addJpeg("file")) // file.jpeg

	fmt.Println(time.Now())
}

func printValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getTo2(input int) (int, int) {
	return 2 * input, 3 * input
}

// 如果定义的命名返回值，则return可以不带参数
// 只要有命名返回值，必须使用()括起来
func getTo3(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return 2, 3
}

// 修改指针变量reply的值
func multiply(a, b int, reply *int) {
	*reply = a * b
}

func min(s ...int) (min int) {
	if len(s) == 0 {
		return 0
	}
	min = s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return
}

func typecheck(values ...interface{}) {
	for _, value := range values {
		switch v := value.(type) {
		case int:
			println("我是int ", v)
		case uint:
			println("我是uint ", v)
		case float32:
			println("我是float32 ", v)
		case float64:
			fmt.Println("我是float64 ", v)
		case string:
			println("我是string ", v)
		case bool:
			println("我是bool ", v)
		default:
			println("我是default ", v)
		}
	}
}

func deferFunc1() {
	println("deferFunc1-start")
	defer deferFunc2()
	// deferFunc2()
	println("deferFunc1-end")
}

func deferFunc2() {
	println("延迟到调用函数结束！")
	a()
	f()
}

// 接收参数的defer语句
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// 有多个defer被注册时，会以逆序执行（类似栈，即先进后出）
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

// 递归函数
// 可递归计算斐波那契数列
func fibonacci(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}

func fv() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

// MakeAddSuffix 一个添加后缀名的工厂函数
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
