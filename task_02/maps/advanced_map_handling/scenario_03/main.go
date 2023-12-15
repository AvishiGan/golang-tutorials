package main

import (
	"fmt"
)

func main() {
	ages := make(map[int]string)

	ages[25] = "Alice"
	ages[30] = "Bob"
	ages[28] = "Charlie"

	fmt.Println(ages)

	fmt.Println(extractUniqueValues(ages))
}

func extractUniqueValues(m map[int]string) []string {

	uniqueValues := make(map[string]bool)

	result := []string{}

	for _, value := range m {

		_, ok := uniqueValues[value]

		if !ok {
			uniqueValues[value] = true
			result = append(result, value)
		}

	}

	return result
}
