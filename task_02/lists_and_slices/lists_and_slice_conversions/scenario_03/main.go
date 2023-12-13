package main

import "fmt"

func main() {

	str := "Golang"
	fmt.Printf("String: %v\n", str)

	// Each byte in the slice represents the ASCII value of a character in the string
	byteSlice := []byte(str)
	fmt.Printf("Byte slice: %v\n", byteSlice)
}
