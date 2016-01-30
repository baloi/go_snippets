package main

import . "fmt"

const Hello = "hello"

var world string

// init() functions are done during program initialization, before main()
func init() {
	Println("In init()")
	world = "world"
	Println(Hello, world)
}

func main() {
	Println("In main()")
}
