package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// 修改字符串
	str := "hello world!"
	c := []byte(str)
	c[0] = 'c'
	s2 := string(c)

	fmt.Println(str) // "hello world!"
	fmt.Println(s2)  // cello world!

	// 获取字符串的子串
	substr := str[2:3]
	fmt.Println(substr)   // l
	fmt.Println(str[1:3]) // el

	// 遍历字符串 for 或者 for-range
	for i := 0; i < len(str); i++ {
		s := str[i]
		fmt.Println(s)            // 打印的是字节数 bytes
		fmt.Println(str[i : i+1]) // 打印的是字符 Unicode
	}

	// for ix, ch := range str {
	// 	...
	// }

	// 获取字符串的字节数
	fmt.Println(len(str))                    // 12
	fmt.Println(len([]int32(str)))           // 12
	fmt.Println(utf8.RuneCountInString(str)) // 12 最快，性能最好

	// 连接字符串
	str1 := "hello "
	str2 := "world!"

	str1 += str2
	fmt.Println(str1) // hello world!

	var buffer bytes.Buffer
	buffer.WriteString(str1)
	fmt.Println(buffer.String()) // hello world!

	// 实际项目中大多使用strings包对字符串进行判断操作
	sl1 := strings.Split(str1, "")
	str2 = strings.Join(sl1, "")
	fmt.Println(sl1)  // [h e l l o   w o r l d !]
	fmt.Println(str2) // hello world!

	// 使用 os 或者 flag 包进行命令行参数的解析
}
