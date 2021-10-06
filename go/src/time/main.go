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
	f.Println(t)                             // 2021-09-30 14:46:26.6407723 +0800 CST m=+0.003236001
	f.Println(t.Day(), t.Month(), t.Year())  // 30 September 2021
	f.Println(t.Format("02 Jan 2006 15:04")) // 30 Sep 2021 14:46
	f.Println(t.Format(time.ANSIC))          // Thu Sep 30 14:46:26 2021
	f.Println(t.Format(time.RFC1123))        // Thu, 30 Sep 2021 14:46:26 CST

	f.Printf("%02d.%d.%d\n", t.Day(), t.Month(), t.Year())               // 30.9.2021
	f.Printf("%02d.%2d.%4d\n", t.Day(), t.Month(), t.Year())             // 30. 9.2021
	var a, _ = f.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) // 30.09.2021
	f.Println(a)                                                         // 11

	f.Println(t.Date())             // 2021 September 30
	f.Println(time.Duration(86400)) // 86.4µs
	f.Println(t.Location())         // Local
	f.Println(t.UTC())              // 2021-09-30 06:42:01.2913251 +0000 UTC

	week := 60 * 60 * 24 * 7 * 1e9
	f.Println(t.Add(time.Duration(week))) // 2021-10-07 14:46:26.6407723 +0800 CST m=+604800.003236001
	f.Println(time.Duration(week))        // 168h0m0s
}
