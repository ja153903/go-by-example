package main

import "fmt"

type rect struct {
	width, height int 
}

// we declare the receiver type a pointer
// we do this because we might want to change the value
// of the receiver object
// or we want to avoid copying on method calls
func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2 * r.width + 2 * r.height 
}

func main() {
	r := rect{ width: 10, height: 5 }

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r

	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
