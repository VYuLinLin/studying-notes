package main

import (
	"fmt"
	f "fmt"
	"strconv"
	"strings"
)

func main() {
	const str = "你好世界"
	s := "hel" + "lo, "
	s += "world!"
	f.Println(`This is a raw string \n`)   // This is a raw string \n
	f.Println(len("asdf"))                 // 4
	f.Println(len(s))                      // 13
	f.Println(len(str))                    // 12
	f.Println(str[0])                      // 228
	f.Println(str[2])                      // 160
	f.Println(strings.HasPrefix(str, "你")) // true

	indexExample()

	sliceExmple()

	strconvExmple()
}

func indexExample() {
	var str string = "Hi, I'm Marc, Hi."

	f.Printf("The position of \"Marc\" is: ")
	f.Printf("%d\n", strings.Index(str, "Marc")) // 8

	f.Printf("The position of the first instance of \"Hi\" is: ")
	f.Printf("%d\n", strings.Index(str, "Hi")) // 0
	f.Printf("The position of the last instance of \"Hi\" is: ")
	f.Printf("%d\n", strings.LastIndex(str, "Hi")) // 14

	f.Printf("The position of \"Burger\" is: ")
	f.Printf("%d\n", strings.Index(str, "Burger")) // -1

	// 	The position of "Marc" is: 8
	// The position of the first instance of "Hi" is: 0
	// The position of the last instance of "Hi" is: 14
	// The position of "Burger" is: -1
}

func sliceExmple() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)

	// 	Splitted in slice: [The quick brown fox jumps over the lazy dog]
	// The - quick - brown - fox - jumps - over - the - lazy - dog -
	// Splitted in slice: [GO1 The ABC of Go 25]
	// GO1 - The ABC of Go - 25 -
	// sl2 joined by ;: GO1;The ABC of Go;25
}

func strconvExmple() {
	var orig string = "666"
	var an int
	var newS string

	f.Println("int的大小为", strconv.IntSize) // int的大小为 64

	an, err := strconv.Atoi(orig)
	f.Println("整数是：", an) // 整数是： 666

	if err != nil {
		f.Println("此参数不能成功转换：", orig)
	}

	an += 5
	newS = strconv.Itoa(an)
	f.Println("新的字符串是：", newS) // 新的字符串是： 671
}
