package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	// go run flag.go -n=6
	var numCores = flag.Int("n", 2, "number of CPU cores to use")
	fmt.Println(*numCores) // 2
	flag.Parse()
	fmt.Println(*numCores) // 6

	runtime.GOMAXPROCS(*numCores)
}
