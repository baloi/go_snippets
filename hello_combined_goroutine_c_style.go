package main

import (
	. "fmt"
	. "net/http"
)

const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1025"

func main() {
	message := "hello world"
	HandleFunc("/hello", func(w ResponseWriter, r *Request) {
		w.Header().Set("Content-Type", "text/plain")
		Fprintf(w, message)
	})

	go func() {
		ListenAndServe(ADDRESS, nil)
	}()

	go func() {
		ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
	}()

	for {
	}
	// we are using an infinite for loop. This DOES NOT work either.
	// Semantically correct, but the way the __goroutine__s are scheduled:
	// the infinite loop starves the thread scheduler and prevents the other
	// __goroutine__s from running.
}
