package main

import (
	"fmt"
)

func main() {
	info := map[string]interface{}{"name": "Alice", "age": 30, "isStudent": true}

	stringMap := make(map[string]string)

	for key, value := range info {
		stringValue, yes := value.(string)
		if yes {
			stringMap[key] = stringValue
		}
	}
	fmt.Println(stringMap)
}
