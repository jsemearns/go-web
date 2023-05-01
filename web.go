package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "hello, you've requested %s\n", r.URL.Path)
	})

	http.ListenAndServe((":80"), nil)
}
