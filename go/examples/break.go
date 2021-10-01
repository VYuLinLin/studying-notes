package main

import "fmt"

func main() {
	m, n := 2, 2
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == 1 {
				break
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("pause")

loop:
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == 1 {
				break loop
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("done")

	// 0 0
	// 1 0
	// pause
	// 0 0
	// done
}
