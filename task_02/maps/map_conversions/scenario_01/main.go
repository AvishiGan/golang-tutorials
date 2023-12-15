package main

import "fmt"

func main() {
	data := [][2]string{{"a", "apple"}, {"b", "banana"}, {"c", "cherry"}}

	dataMap := make(map[string]string)

	for _, pair := range data {
		dataMap[pair[0]] = pair[1] // Adding a new key-value pair to the map
	}
	fmt.Println(dataMap)
}
