package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	fmt.Println("开始输入您的姓名")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s! \n", firstName, lastName)

	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("格式化后的字符串", f, i, s)

	bufioExample()

	bufioReaderExample()

	writeFileExample()

	fscanExample()
}

var inputReader *bufio.Reader
var inputs string
var err error

func bufioExample() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("请输入")

	inputs, err = inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("读取错误")
	}
	fmt.Println("输入的值为： ", inputs)

	switch inputs {
	case "Philip\r\n":
		fmt.Println("welcome Philip!")
	case "Chris\r\n":
		fmt.Println("welcome Chris!")
	default:
		fmt.Println("welcome here!")
	}
}

func bufioReaderExample() {
	inputFile, inputError := os.Open("input.dat")
	// 如果文件不存在或没有权限打开时，会返回错误
	if inputError != nil {
		fmt.Println("打开inputfile时发生错误", inputError)
		return
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	// fmt.Println(inputReader)
	for {
		inputString, err := inputReader.ReadString('\n')
		fmt.Println("input.dat：", inputString)
		if err == io.EOF {
			return
		}
	}
}

func writeFileExample() {
	inputFile := "products.txt"
	outputFile := "products_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "文件错误：", err)
	}
	fmt.Println(string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error)
	}
}

// 读取数据后，按列排列，空格分隔
func fscanExample() {
	file, err := os.Open("products2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
