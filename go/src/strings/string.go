package main

import "fmt"

func main() {
	s := "搜索时间功能"
	// s[0] = "c" // cannot assign to s[0] (value of type byte)
	c := []byte(s)
	c[0] = 'c'

	fmt.Println(string(c), len(c)) // c��索时间功能 18

	r := []rune(s)
	r[0] = '哈'
	fmt.Println(string(r), len(r)) // 哈索时间功能 6

	str := s[len(s)/2:] + s[:len(s)/2]
	fmt.Println(str)
}
