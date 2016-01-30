package main

import . "fmt"

type Text string

func main() {
	var name Text = "namae"
	var pointer_to_name *Text = &name

	Printf("name = %v\n", name)
	Printf("&name = %v\n", &name)
	Printf("pointer_to_name = %v\n", pointer_to_name)
	Printf("value of pointer_to_name (*pointer_to_name) = %v\n", *pointer_to_name)
}
