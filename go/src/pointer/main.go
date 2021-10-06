package main

import f "fmt"

/***
*Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
*每个内存块（或字）有一个地址，通常用十六进制数表示，如：0x6b0820 或 0xf84001d7f0。
* 在指针变量前面加*，可以返回指针变量存储的内存地址当中的值，这被称为反引用。
* 无法获取到文字或常量的内存地址
 */
func main() {
	var a = 1
	i := 123
	f.Println(&a) // 0xc000010090
	f.Println(&i) // 0xc000010098

	var inP *int   // 声明指针变量
	f.Println(inP) // <nil>

	inP = &i
	f.Println(inP)  // 0xc000010098
	f.Println(&inP) // 0xc000006030
	f.Println(*inP) // 123

	// 内存地址没有变，值变了
	*inP = 321
	f.Println(i, inP, &i) // 321 0xc000010098 0xc000010098
	f.Println(a == *(&a)) // true

	// const b = 2
	// f.Println(&b) // cannot take the address of bgo

	var s = "中国人"
	f.Println(&s) // 0xc00003e1f0

}
