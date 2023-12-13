package main

import "fmt"

func main() {
	scores := map[string]float64{
		"English": 85,
		"Maths":   90,
		"Sinhala": 80,
	}
	fmt.Println(scores)
	fmt.Printf("Length: %v", len(scores))
}
