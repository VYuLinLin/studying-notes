package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	openReadFile()
}

// 打开一个文件并读取
func openReadFile() {
	file, err := os.Open("input.dat")
	if err != nil {
		fmt.Println("文件读取错误: ", err.Error())
	}
	defer file.Close()
	iReader := bufio.NewReader(file)

	for {
		str, err := iReader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println(str)
	}
}

// 通过切片读写文件
func cat(f *file.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, er := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading from %s: %s\n", f.String(), er.String())
			os.Exit(1)
		case nr == 0: // EOF
			return
		case nr > 0:
			if nw, ew := file.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing from %s: %s\n",
					f.String(), ew.String())
			}

		}
	}
}
