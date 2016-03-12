package main

import "fmt"

func main() {
	function_channel()
	simple_channel()
}

func simple_channel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func function_channel() {
	// Create a new channel with make (chan val-type)
	// Channels are typed by the values they convey
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax.
	// Here we send "ping" to the messages channel from a new goroutine
	go func() { messages <- "ping" }()

	// the <-channel syntax receives a value from the channel
	// here we'll receive the "ping" message we sent above and print it out
	msg := <-messages

	fmt.Println(msg)

}
