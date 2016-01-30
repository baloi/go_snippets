package main

import (
	. "fmt"
	. "net/http"
)

const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1025"

func main() {
	message = "hello world"
	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprintf(w, message)
	})

	// The go keyword marks a goroutine which is a lightweight thread scheduled
	// by the go runtime. It takes a functional call and creates a separate
	// thread of execution for it. In here we first launch a goroutine to run
	// the HTTP server then run the HTTPS server in the main flow of execution

	go func() {
		ListenAndServe(ADDRESS, nil)
	}()

	ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)

	// Unfortunately in this code, when the main function returns our
	// program will exit and our web servers will terminate

}
