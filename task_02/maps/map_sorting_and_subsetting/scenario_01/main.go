package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)

	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 28
	ages["Harry"] = 25
	ages["Cedric"] = 30
	ages["Luna"] = 28
	ages["Draco"] = 25

	fmt.Println(ages)

	keys := []string{}

	for key := range ages {

		keys = append(keys, key)
	}

	sort.Strings(keys)

	fmt.Printf("Sorted keys: %v\n", keys)
}
