package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var buf bytes.Buffer
	for {
		if s, ok := getNextString(); ok {
			buf.WriteString(strconv.Itoa(s))
		} else {
			break
		}
	}
	buf.WriteString(" - hello")
	buf.WriteString(",")
	buf.WriteString("world")
	buf.WriteString("!")

	bufStr := buf.String()
	fmt.Println(bufStr) // 12345 - hello,world!

	slice1, slice2 := bufStr[:5], bufStr[5:]
	fmt.Printf("%v, %v", slice1, slice2) // 12345 - hello,world!
}

var count int = 0

func getNextString() (int, bool) {
	count++
	if count >= 6 {
		return 0, false
	}
	return count, true
}
