package main

import (
	f "fmt"
	"reflect"
)

func main() {
	var x float32 = 3.14
	var xt = reflect.TypeOf(x)
	var xv = reflect.ValueOf(x)
	f.Println(xt)             // float32
	f.Println(xv)             // 3.14
	f.Println(xv.Interface()) // 3.14

	f.Println(reflect.TypeOf(xt))  // *reflect.rtype
	f.Println(reflect.ValueOf(xt)) // float32

	f.Println(reflect.TypeOf(xv))  // reflect.Value
	f.Println(reflect.ValueOf(xv)) // <float32 Value>

	f.Println(xt.Kind()) // float32
	f.Println(xv.Kind()) // float32

	f.Println(xv.Kind() == reflect.Float32) // true
	f.Println(xv.Kind() == reflect.Float64) // false

	// if reflect.TypeOf(x) == *float32 {
	// 	f.Println("x的类型是float32") // float32
	// }
	constExample()
}

func constExample() {
	// reflect反射包中，包含的Go语言所有的类型常量
	const (
		Kind = iota
		Bool
		Int
		Int8
		Int16
		Int32
		Int64
		Uint
		Uint8
		Uint16
		Uint32
		Uint64
		Uintptr
		Float32
		Float64
		Complex64
		Complex128
		Array
		Chan
		Func
		Interface
		Map
		Ptr
		Slice
		String
		Struct
		UnsafePointer
	)
	// 2 3 4 5 6 7
	f.Println(
		Bool,
		Int,
		Int8,
		Int16,
		Int32,
		Int64,
	)
}
