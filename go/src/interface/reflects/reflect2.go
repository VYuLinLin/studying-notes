package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 3.14
	v := reflect.ValueOf(x)

	// v.SetFloat(3.1415926) // panic: reflect: reflect.Value.SetFloat using unaddressable value
	fmt.Println("can set", v.CanSet()) // can set false

	v = reflect.ValueOf(&x)
	fmt.Println("type of v:", v.Type()) // type of v: *float64
	fmt.Println("can set", v.CanSet())  // can set false

	v = v.Elem()                       // 必须使用Elem方法寻找到变量的指针地址
	fmt.Println("can set", v.CanSet()) // can set true

	v.SetFloat(3.1415926)

	fmt.Println(v.Interface()) // 3.1415926
	fmt.Println(v)             // 3.1415926
	fmt.Println(x)             // 3.1415926
}
