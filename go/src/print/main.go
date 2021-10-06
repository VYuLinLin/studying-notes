package main

import "fmt"

func main() {
	type T struct {
		a int
		b string
	}
	t := T{77, "septemper"}
	a := []int{1, 2, 3, 4}
	b := a[1:3]
	fmt.Printf("%v %v %v \n", t, a, b)                // {77 septemper} [1 2 3 4] [2 3]
	fmt.Printf("%+v %+v %+v \n", t, a, b)             // {a:77 b:septemper} [1 2 3 4] [2 3]
	var aa, _ = fmt.Printf("%#v %#v %#v \n", t, a, b) // main.T{a:77, b:"septemper"} []int{1, 2, 3, 4} []int{2, 3}
	fmt.Printf("%T %T %T \n", t, a, b)                // main.T []int []int
	fmt.Println(aa)
}
