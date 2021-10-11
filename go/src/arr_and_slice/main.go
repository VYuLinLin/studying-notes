package main

import (
	"bytes"
	"fmt"
	"reflect"
	"unicode/utf8"
)

// 打印时强烈建议引用fmt包
// 内置函数print和println打印复杂类型时要么报错，要么只打印内存地址
func main() {
	var arr1 [5]int // 声明一个长度5，类型为int的空数组，默认值为0

	fmt.Println(arr1) // [0 0 0 0 0]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	slice1 := arr1[1:3]            // 定义切片
	slice2 := slice1[:cap(slice1)] // 扩大切片到最大上限
	slice1[0] = 124
	slice2[1] = 123
	fmt.Println(arr1, slice1, slice2) // [0 124 123 6 8] [124 123] [124 123 6 8]
	fmt.Println("切片slice1 的长度为 = ", len(slice1), "最大长度可达 = ", cap(slice1))
	// 切片slice1 的长度为 =  3 最大长度可达 =  5

	var arr22 = []int{1, 2, 3}[:]
	fmt.Println(arr22) // [1 2 3]

	for i, val := range arr1 {
		fmt.Printf("arr1[%d] = %d\n", i, val)
	}
	// arr1[0] = 0
	// arr1[1] = 2
	// arr1[2] = 4
	// arr1[3] = 6
	// arr1[4] = 8

	arr2 := []string{"a", "b", "c"}

	for i := 0; i < len(arr2); i++ {
		fmt.Printf("arr2[%d] = %s\n", i, arr2[i])
	}
	// arr2[0] = a
	// arr2[1] = b
	// arr2[2] = c

	// Go语言中，数组是一种值类型，所以可以通过new()来创建
	arr := new([2]int)
	arr3 := *arr
	arr3[1] = 22
	arr[0] = 11

	fmt.Println("arr 和 arr3是两个不同的数组", arr, arr3)
	// arr 和 arr3是两个不同的数组 &[11 0] [0 22]

	var ar [3]int
	fp(&ar)
	f(ar)
	// fp(&ar)
	fmt.Println(ar)      // [222 0 0]
	fmt.Println(cap(ar)) // 3

	example1()

	sliceExample()

	bufferExample()

	copyAndAppend()

	stringArrExample()
}

// 传入的参数是对原值的拷贝
func f(a [3]int) {
	a[1] = 111
	fmt.Println(a) // [222 111 0]
}

// 传入的参数是引用指针，修改的值是原值
func fp(a *[3]int) {
	a[0] = 222
	fmt.Println(a) // &[222 0 0]
}

func example1() {
	fmt.Println("--------------------------")
	var arr1 [6]int
	var slice1 []int = arr1[2:5]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice1[%d] = %d\n", i, slice1[i])
	}

	fmt.Println("len(arr1) = ", len(arr1))     // 6
	fmt.Println("len(slice1) = ", len(slice1)) // 3
	fmt.Println("cap(slice1) = ", cap(slice1)) // 4

	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice1[%d] = %d\n", i, slice1[i])
	}

	fmt.Println("len(slice1) = ", len(slice1)) // 4
	fmt.Println("cap(slice1) = ", cap(slice1)) // 4
}

func sliceExample() {
	var p *[]int = new([]int)

	fmt.Println(*p, *p == nil) // [] true

	p1 := new([]int)

	fmt.Println(p1) // &[]

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Println(s1, s2) // [112 111 101 109] [101 109]

	s2[1] = 't'
	fmt.Println(s1, s2) // [112 111 101 116] [101 116]

	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7]

	fmt.Println(a) // [5 6]

	a = a[0:4]
	fmt.Println(a) // [5 6 7 8]

}

var (
	s  string = ""
	ok bool
)

func getNextString() (string, bool) {
	fmt.Println(len(s), s, ok)
	if len(s) != 0 && ok {
		ok = false
		return s, ok
	}
	s = "hello,world!"
	ok = true
	return s, ok
}

func bufferExample() {

	fmt.Println("bufferExample", s, ok)
	var buffer bytes.Buffer
	for {
		if s, ok := getNextString(); ok {
			buffer.WriteString(s)
		} else {
			break
		}
		// s, ok = getNextString()
	}
	var buf = buffer.String()
	fmt.Println(buf, ok) // hello,world! false

	slice1, slice2 := buf[:5], buf[len(buf)-1:]
	fmt.Println("slice1 = ", slice1) // hello
	fmt.Println("slice2 = ", slice2) // !

	// reflect.TypeOf() 可以检测任意变量
	var a, b, c = "hh", 2, 0.232
	slice := []int{}
	fmt.Println(reflect.TypeOf(a))                              // string
	fmt.Println(reflect.TypeOf(b))                              // int
	fmt.Println(reflect.TypeOf(c))                              // float64
	fmt.Println(reflect.TypeOf(slice))                          // []int
	fmt.Println(reflect.ValueOf(slice).Type())                  // []int
	fmt.Println(reflect.TypeOf(slice).Kind() == reflect.Slice)  // true
	fmt.Println(reflect.ValueOf(slice).Kind() == reflect.Slice) // true
	// reflect.ValueOf(slice).Type() 等同于 reflect.TypeOf(slice)
	fmt.Println(reflect.ValueOf(slice).Type().Kind() == reflect.Slice) // true

	items := []int{10, 20, 30, 40, 50}
	for i := range items {
		items[i] *= 2
	}
	fmt.Println(items) // [20 40 60 80 100]
}

// 此种方法需要提前知道有几种类型，i.(type)只能在switch中使用，
// 函数没有返回值
func myType(i interface{}) {
	switch i.(type) {
	case string:
	case int:
	}
	return
}

func copyAndAppend() {
	slFrom := []int{1, 2, 3}
	// slFrom1 := []int{11}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)         // [1 2 3 0 0 0 0 0 0 0]
	fmt.Println("已拷贝的数量为", n) // 复制的数量为 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3) // [1 2 3 4 5 6]
}

func stringArrExample() {
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c\n", i, c)
	}

	st := "z界\u754c\ufffd"
	for i, c := range st {
		fmt.Printf("%d:%c\n", i, c)
	}

	c := []int32(s)
	r := []rune(s)
	fmt.Println(c, r, s) // [255 30028] [255 30028] ÿ界
	fmt.Println(st[0:4])

	l := utf8.RuneCountInString(st)
	fmt.Println(l, len([]int32(st)), len(st)) // 4 4 10

	var b []byte
	newS := append(b, s...)
	fmt.Println(newS) // [195 191 231 149 140]

}
