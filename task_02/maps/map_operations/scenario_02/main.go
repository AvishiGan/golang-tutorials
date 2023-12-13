package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 28,
	}

	for key, value := range ages {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
