package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)

	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 28

	info := map[string]int{
		"Harry":  100,
		"Cedric": 50,
		"Luna":   60,
		"Draco":  40,
	}

	fmt.Println(ages)

	fmt.Println(info)

	fmt.Println(mergeMaps(ages, info))
}

func mergeMaps(map1, map2 map[string]int) map[string]int {

	mergedMap := make(map[string]int)

	for k, v := range map1 {
		mergedMap[k] = v
	}

	for k, v := range map2 {
		mergedMap[k] = v
	}

	return mergedMap
}
