package main

import (
	f "fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	int16to32()

	f.Println(Uint8FromInt(10)) // 10 <nil>

	f.Println(IntFromFloat64(6.49)) // 7

	ufunc()

	customRandom()

	byteFn()
}

// int类型转换
func int16to32() {
	var (
		n int16 = 34
		m int32
	)

	m = int32(n)

	f.Printf("32 bit int is: %d\n", m) // 32 bit int is: 34
	f.Printf("16 bit int is: %d\n", n) // 16 bit int is: 34
}

// Uint8FromInt int型转换为int8
func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}

	return 0, f.Errorf("%d 超出了uint8的范围", n)
}

// IntFromFloat64 float64转换为int
func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(f.Sprintf("%g 超出了int32的范围", x))
}

// ByteSize 字节大小
type ByteSize float64

// 使用位左移与iota计数，声明存储单位的常量枚举
const (
	_           = iota // 通过赋值给空白标识符来忽略值
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// 运算/通用函数
func ufunc() {
	f.Println("9 / 4 = ", 9/4) // 9 / 4 =  2
	f.Println("9 % 4 = ", 9%4) // 9 % 4 =  1
	// f.Println("9 % 0 = ", 9/0) // 9 % 4 =  1
}

// 随机数
func customRandom() {
	for i := 0; i < 10; i++ {
		a := rand.Int() // 默认生成20位以内的数字
		f.Println(a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(8) // 生成8以内的数字
		f.Println(r)
	}
	f.Println()
	// 根据当前时间生成随机浮点数
	times := int64(time.Now().Nanosecond())
	rand.Seed(times)
	for i := 0; i < 10; i++ {
		f.Println(100 * rand.Float32())
	}
}

func byteFn() {
	var cha byte = 65         // ASCII
	var ch8 byte = '\377'     // 8进制
	var ch16 byte = '\x41'    // 16进制
	f.Println(cha, ch8, ch16) // 65 255 65

	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	f.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	f.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	f.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	f.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
}
