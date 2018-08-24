package main

import "fmt"

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no messages received")
	}
	
	// in this case we are not sending the message to the channel
	// we are just assigning it
	msg := "hi"
	
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}


	select {
	case msg := <-messages:
		fmt.Println("received messge", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
