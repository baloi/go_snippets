package main

import . "fmt"

type Text string

func main() {
	var name Text = "Ellie"
	var point_to_name *Text

	point_to_name = &name

	Printf("name = %v stored at %v\n", name, point_to_name)
	Printf("point_to_name references %v\n", *point_to_name)
}
