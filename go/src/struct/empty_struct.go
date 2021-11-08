package main

import (
	"fmt"
	"time"
)

type Set map[string]struct{}

func (s Set) Append(k string) {
	s[k] = struct{}{}
}
func (s Set) Remove(k string) {
	delete(s, k)
}
func (s Set) Exist(k string) bool {
	_, ok := s[k]
	return ok
}

/*
	空结构的使用场景：
	实现方法的接收者
	实现集合类型
	实现空通道
*/

// 空结构的应用——实现只用到key的集合的类型
func main() {
	s := Set{}
	s.Append("辣子鸡")
	s.Append("浓汤鸡")
	s.Append("叫花鸡")
	s.Remove("浓汤鸡")

	fmt.Println(s)              // map[叫花鸡:{} 辣子鸡:{}]
	fmt.Println(s.Exist("叫花鸡")) // true
	fmt.Println(s.Exist("浓汤鸡")) // false

	ch := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		close(ch)
	}()
	fmt.Println("咱们老百姓啊")
	a, b := <-ch
	fmt.Println("今天不高兴啊", a, b) // 今天不高兴啊 {} false
}
