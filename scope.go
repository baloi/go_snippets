package main

import . "fmt"

const Hello = "Hello"

var world = "world"

func main() {
	// note that using ":=" instead of "=" will make "world" in main a local
	// variable taking it's value for "world" global concatenated with "!"
	// as shown in the right side
	world := world + "!"
	// The ":=" operator makes an assignment declaration in which the type of
	// expression is inferred from the type of the value being assigned.
	Println(Hello, world)
}
