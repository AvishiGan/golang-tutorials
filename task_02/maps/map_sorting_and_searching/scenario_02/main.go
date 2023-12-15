package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 28,
	}

	if age, exists := ages["Bob"]; exists {
		fmt.Printf("Bob is %d years old", age)
	} else {
		fmt.Println("Not found")
	}
}
