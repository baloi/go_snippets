package main

import (
	"bufio"
	. "fmt"
	"net"
)

func main() {
	// get connection: dial to 1024 and check if error is nil
	if connection, e := net.Dial("tcp", ":1024"); e == nil {
		// defer closing connection
		defer connection.Close()
		// bufio.NewReader  ReadString \n , check if error is nil
		if text, e := bufio.NewReader(connection).ReadString('\n'); e == nil {
			// print the text that NewReader read
			Print(text)
		}
	}
}
