package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// example
// go run main.go output.dat target.txt
func main() {
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	// fmt.Println(os.Args)
	fmt.Println("Good Morning", who)

	// flagExample()

	// bufferFlagExample()

	sliceReadWrite()

	// interfaceWriteExample()
}

var NewLine = flag.Bool("n", false, "print newline")

const (
	Space   = " "
	Newline = "\n"
)

func flagExample() {
	flag.PrintDefaults() //   -n    print newline
	flag.Parse()
	s := ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine {
				s += Newline
			}
		}

		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)                     // target.txt
	os.Stdout.WriteString("\nflagExample end\n") // flagExample end
}

func cat(r *bufio.Reader) {
	var i = 0
	for {
		buf, err := r.ReadBytes('\n')
		i++
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%d %s", i, buf)
	}
}

func bufferFlagExample() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:读取错误 %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
		f.Close()
	}
}

// 利用切片读写文件
func cat2(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, err := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "读取错误： %s\n", err.Error)
			os.Exit(1)
		case nr == 0:
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "写入错误： %s\n", ew.Error())
			}
		}
	}
}

func sliceReadWrite() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat2(os.Stdin)
	}
	fmt.Println("参数：", flag.NArg()) // 参数： 1
	b := []byte{'w', 'o', 'r', 'k', '\n'}
	os.Stdout.Write(b) // work
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if f == nil {
			fmt.Fprintf(os.Stderr, "无法打开%s：error %s\n", flag.Arg(i), err.Error())
			os.Exit(1)
		}
		cat2(f)
		f.Close()
	}
}

func interfaceWriteExample() {
	inputFile, _ := os.Open("target.txt")
	outputFile, _ := os.OpenFile("target1.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerErr := inputReader.ReadLine()
		if readerErr == io.EOF {
			fmt.Println("EOF")
			break
		}
		outputString := string(inputString[2:5]) + "\r\n"
		fmt.Println(inputString, outputString)
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	outputWriter.Flush()
	fmt.Println("Conersion done")
}
