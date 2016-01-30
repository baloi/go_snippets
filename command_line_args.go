package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Program name is %s\n", os.Args[0])
	// check if there is a command line argument after program name
	if len(os.Args) > 1 {
		fmt.Printf("argument after program name is %s\n", os.Args[1])
	}
}
