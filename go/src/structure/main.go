package main

import (
	"fmt"
	f "fmt"
	"runtime"
)

var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
	println(prompt) // Enter a digit, e.g. 3 or Ctrl+Z, Enter to quit.
}

func main() {
	f.Println("当前月份：", Season(7))

	f.Println(Abs(-5)) // 5

	if true {
		f.Println("成功条件代码")
	}

	if runtime.GOOS == "windows" {
		f.Println("当前系统开发商为：微软")
	}
	// f.Println("当前系统开发商为：", runtime.GOOS)

	for i := 0; i < 3; i++ {
		// 当前系统开发商为： windows
		f.Println("当前系统开发商为：", runtime.GOOS)
	}

	str := "hello,world!"
	for pos, char := range str {
		f.Printf("Character on position %d is: %c \n", pos, char)
		// f.Println("position=", pos, "character=", char)
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v) // 0
		v = 5
	}
	// 无限循环
	// for i := 0; ; i++ {
	// 	fmt.Println("Value of i is now:", i)
	// }

	// break
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
				print("break之后的代码") // 此代码不会执行
			}
			print(j)
		}
		println("  ")
	}

	labelExample()

	gotoExample()

	gotoExample1()

	i := 0
	for { //since there are no checks, this is an infinite loop
		if i >= 3 {
			break
		}
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}
}

// Abs 获取整型数字的绝对值
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x

}

// 比较两个整型数字的大小
func isGreater(x, y int) bool {
	if x > y {
		return true
	}
	return false
}

// Season 获取
func Season(month int) string {
	switch month {
	case 0, 12:
		return "十二月"
	case 1:
		return "一月"
	case 2:
		return "二月"
	case 3:
		return "三月"
	case 4:
		return "四月"
	case 5:
		return "五月"
	case 6:
		return "六月"
	case 7:
		return "七月"
	case 8:
		return "八月"
	case 9:
		return "九月"
	case 10:
		return "十月"
	case 11:
		return "十一月"
	default:
		return "未知月份"
	}
}

// 标签循环
func labelExample() {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				// 与不使用标签时的break效果一样
				continue LABEL1
			}
			f.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}

// goto 循环
func gotoExample() {
	i := 0
HERE:
	f.Printf("goto%d\n", i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}

func gotoExample1() {
	a := 1
	b := 9
	goto TARGET // compile error
	b += a
TARGET:
	fmt.Printf("a is %v *** b is %v\n", a, b) // a is 1 *** b is 9

}
