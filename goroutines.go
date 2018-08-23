package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	// goroutines allow us to asynchronously run functions
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	
	// we're required to have scanln here
	// or in production we can put it to sleep
	time.Sleep(5000)
	//fmt.Scanln()
	fmt.Println("done")
}
