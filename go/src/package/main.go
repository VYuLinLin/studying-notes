package main

import (
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
)

func main() {
	regexpExample()

	bigExample()
}

// regexp 正则包
func regexpExample() {
	searchIn := "John: 2578.347 william: 4567.23 steve: 56.37"
	reg := "[0-9]+.[0-9]+" // 正则

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(reg, []byte(searchIn)); ok {
		fmt.Println("发现匹配")
	}

	re, _ := regexp.Compile(reg)
	// 将匹配的部分替换为"##.#"
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str) // John: ##.# william: ##.# steve: ##.#

	// 参数为函数时
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2) // John: 5156.69 william: 9134.46 steve: 112.74

}

func bigExample() {
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(2)
	ip.Mul(io, ip).Div(big.NewInt(666), big.NewInt(6))
	fmt.Println("Big Int: ", ip) // Big Int:  111
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Println("Big Max Int: ", im, len(fmt.Sprint(im))) // Big Max Int:  9223372036854775807 19
	fmt.Println("Big Int: ", ip)                          // Big Int:  43492122561469640008497075573153004

	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1946, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Println("Big Rat: ", rq) // Big Rat:  -17953/54768
}

func init() {
	fmt.Println("init one")
}
func init() {
	fmt.Println("init two")
}
