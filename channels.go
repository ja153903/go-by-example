package main

import "fmt"

func main() {
	
	messages := make(chan string)
	
	// we have a goroutine that indefinitely
	// sends ping to messages
	go func() { messages <- "ping" }()
	
	// then we have msg which receives the pings
	msg := <-messages 
	fmt.Println(msg)
}
