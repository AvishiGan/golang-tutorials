package main

import "fmt"

func main() {
	ages := map[string]float64{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 28,
	}
	fmt.Println(ages)

	total := totalMarks(ages)

	fmt.Printf("Total marks: %.2f\n", total)
}

func totalMarks(scores map[string]float64) float64 {
	var total float64

	for _, value := range scores {
		total += value
	}

	return total
}
