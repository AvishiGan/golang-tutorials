package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)

	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 28

	fmt.Println(ages)

	submap := make(map[string]int)

	for key, value := range ages {

		if key == "Bob" || key == "Charlie" {
			submap[key] = value
		}
	}

	fmt.Printf("Submap: %v", submap)
}
