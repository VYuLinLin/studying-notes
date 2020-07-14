// hello world包的简介描述
package main

import (
	f "fmt" // 别名
	"math"
	"runtime"
	"strings"
)

// Pi sss
var Pi float64

// init为初始化函数
func init() {
	Pi = 4 * math.Atan(1)
	f.Print("I'm a init function\n")
	fn()
}

// main是每个文件中必须有的函数且在init函数后执行
func main() {
	// var _ = "234"
	// f.Print(_) // 报错 cannot use _ as value
	var t = "hello, world"
	f.Println(t)

	str0 := "hello world"
	str1 := "hello " +
		`world1`
	str2 := []string{"hello", "world"}
	str3 := strings.Join(str2, ",")

	f.Println(str0[3])           // 108
	f.Println(str0[3:5])         // lo
	f.Println(str0[3:len(str0)]) // lo world
	f.Println(str1)              // hello world1
	f.Println(str2, len(str2))   // [hello world] 2
	f.Println(str3)              // helflo,world

	var goos string = runtime.GOOS
	f.Printf("操作系统类型:%s\n", goos) // windows
	// path := os.Getenv("PATH")
	// f.Printf("go 路径%s\n", path) // 系统中所有环境变量中的值

	// var a int = 6
	var b int32
	// b = b + a // 报错
	b = b + 5
	f.Println(b) // 5 能编译成功是因为5是常量

}

// fn 为私有函数
func fn() {
	f.Println("I'm a custom function")
	F1()
}

// F1 为导出函数
func F1() {
	j := 32 // 变量的简写形式，只能在函数体内使用
	f.Println("I'm a custom function one")
	f.Println(a, b, c) // 0 1 2
	f.Println(n, j)    // 2 32
	// 同一类型的多个变量声明
	var aa, bb, cc int
	// 多变量在同一行赋值，此方式只能在函数体内使用，不能在全局作用域使用
	aa, bb, cc = 55, 77, 22
	f.Println(aa, bb, cc) // 55 77 22
}

// 并行赋值的方式定义枚举值常量
// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )

// 简写
const (
	a = iota
	b
	c
	d
	e
)

// 声明变量
var n int64 = 2
