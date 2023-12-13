package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 28,
	}

	if _, exists := ages["Alice"]; exists {
		fmt.Println("Exists")
	} else {
		fmt.Println("Not found")
	}
}
