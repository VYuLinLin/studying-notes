package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("内部回调", req.URL.Path[1:])
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", req.URL.Path[1:], "hello, world!")
}

func main() {
	// 写法一
	// http.HandleFunc("/", HelloServer)
	// err := http.ListenAndServe("localhost:5001", nil)

	// 写法二
	fmt.Println("starting serve")

	err := http.ListenAndServe("localhost:5000", http.HandlerFunc(HelloServer))
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
	fmt.Println("ending serve")
}
