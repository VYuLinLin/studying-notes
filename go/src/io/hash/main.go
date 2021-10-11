package main

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"log"
)

func main() {
	hasher := sha1.New()
	io.WriteString(hasher, "test")
	b := []byte{}
	fmt.Printf("结果：%x\n", hasher.Sum(b))
	// 下面两种效果相同
	fmt.Printf("结果：%d\n", hasher.Sum(b))

	hasher.Reset()
	data := []byte("we shall overcome!")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("hash write error: %v / %v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("结果： %x\n", checksum)

	e := errors.New("犯错了")
	fmt.Println(e.Error())
}
