package main

import "fmt"

func main() {
	ages := make(map[string]int)

	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 28
	ages["Harry"] = 25
	ages["Cedric"] = 30
	ages["Luna"] = 25
	ages["Draco"] = 25

	fmt.Println(ages)

	fmt.Printf("There are %d individuals who are 25 years old\n", countValues(ages, 25))
}

func countValues(m map[string]int, value int) int {

	count := 0

	for _, v := range m {
		if v == value {
			count++
		}
	}

	return count
}
