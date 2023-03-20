package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello, you requested %s</h1>", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
