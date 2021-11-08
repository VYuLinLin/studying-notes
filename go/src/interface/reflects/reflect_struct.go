package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type NotknownType struct {
	S1, S2, S3 string
}

func (n NotknownType) String() string {
	return n.S1 + " - " + n.S2 + " - " + n.S3
}

// variable to investigate:
var Secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(&Secret).Elem() // <main.NotknownType Value>
	typ := reflect.TypeOf(Secret)            // main.NotknownType
	// alternative:
	//typ := value.Type()
	fmt.Println(typ) // main.NotknownType

	knd := value.Kind()
	fmt.Println(knd)                                // interface
	fmt.Println("I can set value?", value.CanSet()) // I can set value? true

	v := value.Elem()
	fmt.Println(v.Kind())                       // struct
	fmt.Println("I can set value?", v.CanSet()) // I can set value? false
	// iterate through the fields of the struct:
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, v.Field(i))
		// 创建临时结构类型
		tmp := reflect.New(v.Type()).Elem()
		tmp.Set(v)

		// fmt.Println(tmp.Type()) // main.NotknownType

		// panic: reflect: reflect.Value.SetString using unaddressable value
		// v.Field(i).SetString("C#")

		tmp.Field(i).SetString("C#" + strconv.Itoa(i))
		value.Set(tmp)
		v = value.Elem()
	}

	// call the first method, which is String():
	results := v.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]

	fmt.Println(value) // [Ada - Go - Oberon]

	fmt.Println("=============================================")

	struct1 := NotknownType{"Go", "is", "hard"}
	ModifyIt(&struct1, "S3", "GoLang")
	fmt.Println("new struct1:", struct1) // new struct1: Go - is - GoLang

	struct2 := NotknownType{"JS", "or", "bad"}
	ModifyIt2(&struct2, "S3", "GoLang")
	fmt.Println("new struct2:", struct2) // new struct2: JS - or - GoLang
}

func ModifyIt(p interface{}, fieldName string, val interface{}) interface{} {
	value := reflect.ValueOf(p)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	// newVal := reflect.ValueOf(&val).Elem().Interface().(string)
	// or
	newVal := val.(string)

	value.FieldByName(fieldName).SetString(newVal)

	return p
}

func ModifyIt2(p interface{}, fieldName string, val interface{}) interface{} {
	value := reflect.ValueOf(p)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	var newVal string
	if str, ok := val.(string); ok {
		newVal = str
	}
	// 创建临时类型变量
	// fmt.Println("value type: ", value.Type()) // value type:  main.NotknownType
	tmp := reflect.New(value.Type()).Elem()
	tmp.Set(value)
	tmp.FieldByName(fieldName).SetString(newVal)
	value.Set(tmp)

	return p
}
