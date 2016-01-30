package main

import "fmt"

type Hello struct{}

func (h Hello) String() string {
	return "Hello"
}

type World struct{}

func (w *World) String() string {
	return "world"
}

type Message struct {
	X fmt.Stringer
	Y fmt.Stringer
}

func (v Message) IsGreeting() (ok bool) {
	// now, take note of the expression: v.X.(*Hello)
	// A pointer to type Hello? or a type of Hello?

	// The assertion tells us whether one of __Messages data fields
	// contains a value of concrete type Hello
	if _, ok = v.X.(*Hello); !ok {
		_, ok = v.Y.(*Hello)
	}
	return
}

func main() {
	// observe &Message{] is how we created the struct here
	m := &Message{}
	fmt.Printf("m.IsGreeting() --> %v\n", m.IsGreeting())

	m.X = &Hello{}
	fmt.Printf("m.X = &Hello{} --> %v\n", m.IsGreeting())

	m.X = new(Hello)
	fmt.Printf("m.X = new(Hello) --> %v\n", m.IsGreeting())

	m.Y = new(World)
	fmt.Printf("m.Y = new(World) --> %v\n", m.IsGreeting())

	m.X = m.Y
	fmt.Printf("m.X = m.Y --> %v\n", m.IsGreeting())

	m = &Message{X: new(World), Y: new(Hello)}

	fmt.Printf("m = &Message{X: new(World), Y: new(Hello)} --> %v\n", m.IsGreeting())

	m.X, m.Y = m.Y, m.X
	fmt.Printf("m.X, m.Y = m.Y, m.X --> %v\n", m.IsGreeting())
}
