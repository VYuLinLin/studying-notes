package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Println("打开或创建文件时出错")
		return
	}
	defer outputFile.Close() // 在此函数执行完或return后执行

	outputWriter := bufio.NewWriter(outputFile)
	str := "hello world!\n"

	for i := 0; i < 10; i++ {
		outputWriter.WriteString(str)
	}
	outputWriter.Flush()

	outputFile, outputError = os.OpenFile("output.dat", os.O_RDONLY, 0666)
	fmt.Println(bufio.NewReader(outputFile).ReadString('\n'))

	fileCopyExample()
}

func fileCopyExample() {
	copyFile("target.txt", "source.txt")
	fmt.Println("Copy done!")
}

func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
