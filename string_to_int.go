package main

import "fmt"
import "strconv"

func main() {
	i, _ := strconv.Atoi("22")
	s := strconv.Itoa(22)

	fmt.Printf("Integer = %d\n", i)
	fmt.Printf("String = %s\n", s)
}
