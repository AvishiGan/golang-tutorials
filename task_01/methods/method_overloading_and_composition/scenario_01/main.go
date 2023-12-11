package main

import (
	"fmt"
	"math"
)

type Shape interface {
	calculate() float64
}

type Rectangle struct {
	height, width float64
}

// Area
func (r Rectangle) calculate() float64 {
	return r.height * r.width
}

type Circle struct {
	radius float64
}

// Perimeter
func (c Circle) calculate() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := Rectangle{height: 5, width: 7}
	c := Circle{radius: 5}

	fmt.Printf("Area of rectangle: %.2f\n", r.calculate())
	fmt.Printf("Perimeter of circle: %.2f\n", c.calculate())
}
