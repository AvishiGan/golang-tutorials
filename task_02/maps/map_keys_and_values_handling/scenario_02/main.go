package main

import "fmt"

func main() {
	scores := map[string]int{
		"English": 85,
		"Maths":   90,
		"Sinhala": 80,
	}
	fmt.Println(scores)

	values := []int{}

	for _, value := range scores {

		values = append(values, value)
	}

	fmt.Printf("Values: %v\n", values)
}
