package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 28,
	}
	fmt.Println(ages)

	fmt.Printf("People who are older than 25: %v\n", filterMap(ages, 25))
}

func filterMap(m map[string]int, threshold int) map[string]int { // threshold -> a certain level or point that must be exceeded

	filteredMap := make(map[string]int)

	for key, value := range m {

		if value > threshold {

			filteredMap[key] = value
		}
	}

	return filteredMap
}
