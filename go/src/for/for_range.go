package main

import "fmt"

func main() {
	var items = []int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}
	fmt.Printf("%v \n", items) // [10 20 30 40 50]

	for i := range items {
		items[i] *= 2
	}
	fmt.Printf("changed %v", items)
}
