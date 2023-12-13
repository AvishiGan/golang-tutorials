package main

import "fmt"

func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	fmt.Printf("Length of the slice: %d\n", len(names)) // Number of elements it currently contains
	fmt.Printf("Capacity of the slice: %d", cap(names)) // Total amount of storage allocated for the slice
}
