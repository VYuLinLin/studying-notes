package main

import "fmt"

func main() {

LABEL1:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if j == 2 {
				break LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
	fmt.Println("===")
LABEL2:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if j == 2 {
				continue LABEL2
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}
