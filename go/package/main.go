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
	fmt.Println(str)

	// 参数为函数时
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)

}

func bigExample() {
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Println("Big Int: ", ip)

	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1946, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Println("Big Rat: ", rq)
}
