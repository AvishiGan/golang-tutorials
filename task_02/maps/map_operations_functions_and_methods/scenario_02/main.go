package main

import "fmt"

type Details struct {
	info map[string]string
}

func (details Details) displayInfo() {
	for key, value := range details.info {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}

func main() {
	details := Details{
		info: map[string]string{
			"Harry":  "Gryffindor",
			"Cedric": "Hufflepuff",
			"Luna":   "Ravenclaw",
			"Draco":  "Slytherin",
		},
	}
	details.displayInfo()
}
