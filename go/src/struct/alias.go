package main

import (
	"fmt"
	"time"
)

// example one
type type0 struct {
	p int
}
type myTime struct {
	time.Time
	type0
}

func (t myTime) first3Chars() string {
	return t.Time.String()[:4]
}
func (t myTime) getWall() int {
	return t.p
}
func timeExample() {

	m := myTime{Time: time.Now()}

	fmt.Println(m.p)             // 2020-07-31 14:36:40.017164 +0800 CST m=+0.020919201
	fmt.Println(m.String())      // 2020-07-31 14:36:40.017164 +0800 CST m=+0.020919201
	fmt.Println(m.first3Chars()) // 2020
}

// example two
type myTime1 time.Time   // 类型定义
type myTime2 = time.Time // 类型别名
func (t myTime1) first4Chars() string {
	return fmt.Sprintf("%v", t)
}

// 非本地包的类型别名，不能定义方法
// func (t myTime2) first4Chars() string {
// 	return "alias myTime2"
// }
func main() {
	timeExample()
	var a myTime1
	fmt.Printf("a Type: %T \n", a)                    // a Type: main.myTime1
	fmt.Printf("a Type Func: %v \n", a.first4Chars()) // a Type Func: myTime1

	var b myTime2
	fmt.Printf("b Type: %T \n", b) // b Type: time.Time
}
