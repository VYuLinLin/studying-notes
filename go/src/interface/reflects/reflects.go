package main

import (
	f "fmt"
	"reflect"
)

func main() {
	var x float32 = 3.14
	var xt = reflect.TypeOf(x)
	var xv = reflect.ValueOf(x)
	f.Println(xt) // float32
	f.Println(xv) // 3.14
	// f.Println(xv == 3.14) // invalid operation: xv == 3.14 (mismatched types reflect.Value and float64)
	f.Println(xv.Interface())         // 3.14
	f.Println(xv.Interface() == 3.14) // false 因为3.14默认类型是float64所以返回false，可以使用 float32(3.14) 转换类型
	f.Println(reflect.TypeOf(3.14))   // float64

	f.Println(reflect.TypeOf(xt))  // *reflect.rtype
	f.Println(reflect.ValueOf(xt)) // float32

	f.Println(reflect.TypeOf(xv))  // reflect.Value
	f.Println(reflect.ValueOf(xv)) // <float32 Value>

	// Kind 总是返回reflect的底层类型
	f.Println(xt.Kind()) // float32
	f.Println(xv.Kind()) // float32

	f.Println(xv.Kind() == reflect.Float32) // true
	f.Println(xv.Kind() == reflect.Float64) // false

	f.Println(xv.Float())  // 3.140000104904175
	f.Println(xv.String()) // <float32 Value>
	// f.Println(xv.Int())    // panic: reflect: call of reflect.Value.Int on float32 Value

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
