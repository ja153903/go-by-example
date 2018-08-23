package main

import "fmt"

func main() {
	// when you go over a string
	// you actually get back the rune
	// rather than the character
	// there is a need to convert it
	// if that's what you want to do
	for i, c := range "go" {
		fmt.Println(i, c)
		fmt.Println(string(c))
	}
	
	kvs := map[string]int{
		"jaime": 23,
		"jackson": 24,
	}

	for key, value := range kvs {
		fmt.Println(key, value)
	}
}
