package main

import "fmt"

func main() {
	x, y := 10, 20
	swap(&x, &y)
}

func swap(x, y *int) {
	fmt.Printf("Before swap x: %d, y: %d\n", *x, *y)
	*x, *y = *y, *x
	fmt.Printf("After swap x: %d, y: %d\n", *x, *y)
}
