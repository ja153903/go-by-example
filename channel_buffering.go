package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	// fmt.Println(<-messages) another call would cause an error because
	// the buffer is empty
}
