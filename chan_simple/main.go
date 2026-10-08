package main

import "fmt"

func main() {
	// Create a channel that can carry strings between goroutines.
	message := make(chan string)

	// Start a new goroutine. It sends a value into the channel.
	go func() {
		message <- "hello from another goroutine"
	}()

	// Receive the value from the channel.
	// This waits until the goroutine above has sent its message.
	received := <-message

	// Print the value that was received.
	fmt.Println(received)
}
