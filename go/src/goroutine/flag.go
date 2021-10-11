package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	var numCores = flag.Int("n", 2, "number of CPU cores to use")
	fmt.Println(*numCores)

	flag.Parse()
	runtime.GOMAXPROCS(*numCores)
}
