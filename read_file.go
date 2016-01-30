package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	input := flag.String("f", "", "input file")
	flag.Parse()

	f, err := os.Open(*input)

	if err != nil {
		log.Fatal(err)
	}

	//for scanner := bufio.NewScanner(f); scanner.Scan(); fmt.Print("q> ") {
	for scanner := bufio.NewScanner(f); scanner.Scan(); {
		//fmt.Println(">>>", scanner.Bytes())
		fmt.Println(">>>", scanner.Text())
	}
}
