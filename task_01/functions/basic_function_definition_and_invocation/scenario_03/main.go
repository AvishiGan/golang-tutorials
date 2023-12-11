package main

import "fmt"

func main() {
	calculateCircleArea(5)
}

func calculateCircleArea(radius float64) float64 {
	area := 3.14 * radius * radius
	fmt.Println(area)
	return 0
}
