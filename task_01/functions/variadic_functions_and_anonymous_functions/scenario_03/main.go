package main

import "fmt"

func main() {
	concatenate("abcdefghi", "jklmnopqrs", "tuvwxyz")
}

func concatenate(strings ...string) string {
	var result string

	for _, str := range strings {
		result += str
	}
	fmt.Println(result)
	return (result)
}
