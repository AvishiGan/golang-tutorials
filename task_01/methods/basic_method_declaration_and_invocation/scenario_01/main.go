package main

import "fmt"

type Rectangle struct {
	length, width float64
}

func (r Rectangle) calculateArea() float64 {
	area := r.length * r.width
	return area
}

func main() {
	r := Rectangle{length: 5, width: 4}
	fmt.Printf("Area of rectangle: %.2f\n", r.calculateArea())
}
