package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 28,
	}
	fmt.Println(getAverageAge(ages))
}

func getAverageAge(ages map[string]int) float64 {

	var sum int
	var number int

	for _, value := range ages {
		sum += value
		number++
	}

	average := sum / number

	return (float64(average))
}
