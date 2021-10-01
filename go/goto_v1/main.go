package main

import (
	"fmt"
	"net/http"

	"goto_v1/store"
)

const AddForm = `
	<form method="POST" action="/add">
		URL: <input type="text" name="url" />
		<input type="submit" value="Add"/>
	</form>
`

var stores = store.NewURLStore()

func main() {
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":8888", nil)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := stores.Get(key)
	if url == "" {
		http.NotFound(w, r) // 404
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, AddForm)
		return
	}
	key := stores.Put(url)
	fmt.Fprintf(w, "http://localhost:8888/%s", key)
}
