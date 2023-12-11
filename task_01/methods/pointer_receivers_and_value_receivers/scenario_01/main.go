package main

import "fmt"

type Rectangle struct {
	length, width float64
}

func (r *Rectangle) increaseLengthByTwice() {
	r.length *= 2
}

func main() {
	r := Rectangle{length: 5, width: 4}
	r.increaseLengthByTwice()
	fmt.Printf("Updated length of rectangle: %.2f\n", r.length)
}
