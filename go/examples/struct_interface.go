package main

import (
	"fmt"
	"log"
)

// 创建结构体
type struct1 struct {
	field1 string
	field2 int
}

// 创建接口
type Stringer interface {
	fn()
}

func main() {
	// 初始化
	ms1 := new(struct1)
	// or
	ms2 := &struct1{"one", 2}
	fmt.Println(ms1) // &{ 0}
	fmt.Println(ms2) // &{one 2}

	// 当结构体的命名以大写字母开头时,该结构体在包外可见

	// 推荐使用构建函数初始化结构体
	ms3 := Newstruct1("first", 2)
	fmt.Println(ms3) // &{first 2}

}

func Newstruct1(s string, n int) *struct1 {
	return &struct1{s, n}
}

// 通过内建函数recover 终止 panic 过程
func protect(g func()) {
	defer func() {
		log.Println("done")
		// Println executes normally even if there is a panic
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}
