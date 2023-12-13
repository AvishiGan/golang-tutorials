package main

import (
	"fmt"
	"reflect"
)

func main() {
	data := []interface{}{1, "Hello", true}
	fmt.Printf("Slice: %v\n", data)

	stringSlice := []string{}

	for _, i := range data {
		// Check if the value is a string
		if reflect.TypeOf(i).Kind() == reflect.String {

			stringSlice = append(stringSlice, i.(string))
		}
	}
	fmt.Printf("Slice only with strings: %v\n", stringSlice)
}
