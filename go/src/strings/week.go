package main

import (
	"fmt"
)

type Day int

var weeks = [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

const (
	Mon Day = iota
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

func (d Day) String() string {
	return fmt.Sprintf("%s", weeks[d])
}
func main() {
	fmt.Println(Mon) // Monday
	fmt.Println(Tue)
	fmt.Println(Wed)
	fmt.Println(Thu)
	fmt.Println(Fri)
	fmt.Println(Sat)
	fmt.Println(Sun) // Sunday
}
