package main

import "fmt"

type Message struct {
	X string
	y *string
}

func (v Message) Print() {
	if v.y != nil {
		fmt.Println(v.X, *v.y)
	} else {
		fmt.Println(v.X)
	}
}

func (v *Message) Store(x, y string) {
	v.X = x
	v.y = &y
}

func main() {
	m := Message{}
	m.Print()
	yy := "world"
	//m.Store("Hello", "world")
	m.Store("Hello", yy)
	m.Print()

	m.X = "lohe"
	m.Print()

	yy = "ehol"
	m.Print()
}
