package main

import (
	. "fmt"
	"net/http"
)

const SECURE_ADDRESS = ":1025"

func main() {
	message := "hello world ssl\n"
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprintf(w, message)
	})

	http.ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
}
