package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readInput := bufio.NewReader(os.Stdin)
	var val float64
	var EOF = "q\r\n"
	for {
		input1, err := readInput.ReadString('\n')
		if err != nil || input1 == EOF {
			fmt.Println(err)
			break
		}
		operator, err := readInput.ReadString('\n')
		if err != nil || input1 == EOF {
			fmt.Println(err)
			break
		}
		input2, err := readInput.ReadString('\n')
		if err != nil || input1 == EOF {
			fmt.Println(err)
			break
		}
		val1, err := strconv.ParseFloat(input1[:len(input1)-2], 64)
		if err != nil {
			fmt.Println(err)
			break
		}
		var val2 float64
		fmt.Sscanf(input2, "%v\r\n", &val2) // 过滤空格
		// or
		// val2, err := strconv.ParseFloat(input2[:len(input2)-2], 64)
		// if err != nil {
		// 	fmt.Println(err)
		// 	break
		// }
		fmt.Sscanf(operator, "%s", &operator) // 过滤空格
		val = 0
		switch operator {
		case "+":
			val = val1 + val2
		case "-":
			val = val1 - val2
		case "*":
			val = val1 * val2
		case "/":
			val = val1 / val2
		}
		// fmt.Printf("计算1 = %d \n", val1)
		// fmt.Printf("计算2 = %d \n", val2)
		fmt.Printf("计算结果%v %s %v = %v \n", val1, operator, val2, val)
	}
}
