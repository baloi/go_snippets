package main

import "fmt"

func sup(name string) string {
	return "Sup, " + name
}

func main() {
	fmt.Println(sup("Ewan"))
}
