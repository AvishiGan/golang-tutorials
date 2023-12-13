package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3} // An array has a fixed size
	fmt.Printf("Array: %v\n", array)

	slice := array[0:2] // Conversion
	fmt.Printf("Slice: %v", slice)
}
