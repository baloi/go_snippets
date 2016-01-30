package main

import (
	. "fmt"
	"net/http"
)

const MESSAGE = "Hello world\n"
const ADDRESS = ":1024"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprintf(w, MESSAGE)
	})

	Printf("Listen and serving...\n")
	http.ListenAndServe(ADDRESS, nil)
}
