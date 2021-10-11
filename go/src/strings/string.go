package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "搜索时间功能"
	// s[0] = "c" // cannot assign to s[0] (value of type byte)
	c := []byte(s)
	c[0] = 'c'

	fmt.Println(string(c), len(c)) // c��索时间功能 18

	// 替换第一个字符
	r := []rune(s)
	r[0] = '哈'
	fmt.Println(string(r), len(r)) // 哈索时间功能 6

	// 前后颠倒
	str := s[len(s)/2:] + s[:len(s)/2]
	fmt.Println(str)

	funcString()
}

type TwoInts struct {
	a int
	b int
}

func funcString() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
}

// 打印时，会默认调用String方法
func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
