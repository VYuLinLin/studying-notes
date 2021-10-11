package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Iface interface {
	String() string
}

type StrInt interface {
	String() int
}

type User struct {
	name string
	age  int
}

func (u *User) String() string {
	return fmt.Sprint(u.name)
}
func main() {
	// 测试一个值是否实现了某个接口
	user := &User{"xiao v", 19}

	v := Stringer(user)
	// 结构类型断言
	if sv, ok := v.(*User); ok {
		fmt.Printf("[struct ok]user name is %s \n", sv.String())
	}
	// 接口类型断言
	if sv, ok := v.(Stringer); ok {
		fmt.Printf("[interface ok]user name is %s \n", sv.String())
	}

	// 检查结构类型User是否实现了Stringer接口
	var _ Stringer = (*User)(nil) // （此方法不会分配内存）
	// or
	// var _ Stringer = User{} // 此方法会分配内存

	// 接口到接口
	// fmt.Println(v.(StrInt).String()) // panic: interface conversion: *main.User is not main.StrInt: missing method String
	fmt.Println(v.(Iface).String()) // xiao v

	testInterface(v)
}

func testInterface(x Stringer) {
	fmt.Println(x.(Iface).String()) // xiao v
}
