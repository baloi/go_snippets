package main

import (
	. "fmt"
	"net"
)

func main() {
	// create listener: if net.Listen to tcp @ :1024 port is not nil and error == nil
	if listener, e := net.Listen("tcp", ":1024"); e == nil {
		// go into a loop
		for {
			// accept connection and check if error is nil
			if connection, e := listener.Accept(); e == nil {
				// create a goroutine with net.Conn as param (passing connection at end of function def
				go func(c net.Conn) {
					// defer closing Conn
					defer c.Close()
					// Fprintf "hello or something" to Conn object
					Fprintf(c, "hello there\n")
					Printf("sent string to client\n")
				}(connection) // notice that "connection object gets passed here..."
			}
		}
	}
}
