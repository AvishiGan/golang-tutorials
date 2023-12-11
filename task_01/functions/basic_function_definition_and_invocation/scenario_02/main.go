package main

import "fmt"

func main() {
	add(1, 2)
}

func add(a, b int) int {
	sum := a + b
	fmt.Println(sum)
	return 0
}
