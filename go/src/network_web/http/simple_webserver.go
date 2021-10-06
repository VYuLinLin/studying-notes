package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

const form = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="submit" value="submit" />
		</form>
	</body></html>
`

func main() {
	http.HandleFunc("/", logPanics(defaultHandleFunc))
	http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}

func defaultHandleFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	panic(errors.New("test1 error"))
}

// 一个简单的get请求
func SimpleServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>hello, world!</h1>")
}

func FormServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Test2 callback")
	w.Header().Set("Content-Type", "text/html")
	switch req.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, req.FormValue("in"))
	}
}

func logPanics(f HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", req.RemoteAddr, x)
				// 给页面一个错误信息, 如下示例返回一个500
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		f(w, req)
	}
}
