package main

import (
	f "fmt"
	"time"
)

func main() {
	timeExample()
}

func timeExample() {
	t := time.Now()
	f.Println(t)                             // 2020-07-13 11:54:09.1173123 +0800 CST m=+0.001995701
	f.Println(t.Day(), t.Month(), t.Year())  // 13 July 2020
	f.Println(t.Format("02 Jan 2006 15:04")) // 13 Jul 2020 13:34
	f.Println(t.Format(time.ANSIC))          // Mon Jul 13 13:39:42 2020
	f.Println(t.Format(time.RFC1123))        // Mon, 13 Jul 2020 13:39:42 CST
}
