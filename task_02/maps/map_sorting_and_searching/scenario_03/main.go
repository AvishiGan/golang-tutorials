package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 28,
	}
	fmt.Println(ages)

	fmt.Printf("Reversed: %v", reverseMap(ages))
}

func reverseMap(m map[string]int) map[int]string {

	reversedMap := make(map[int]string)

	for key, value := range m {
		reversedMap[value] = key
	}
	return reversedMap
}
