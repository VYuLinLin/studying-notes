package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	var (
		host          = "www.apache.org"
		port          = "80"
		remote        = host + ":" + port
		msg    string = "get / \n"
		data          = make([]uint8, 4096)
		read          = true
		count         = 0
	)
	// 创建一个socket
	con, err := net.Dial("tcp", remote)
	// 发送一个http get 请求
	io.WriteString(con, msg)
	// 读取服务起的响应
	for read {
		count, err = con.Read(data)
		read = (err == nil)
		fmt.Println(string(data[0:count]))
	}
	con.Close()
}
