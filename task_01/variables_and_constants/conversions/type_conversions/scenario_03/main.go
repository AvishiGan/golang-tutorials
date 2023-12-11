package main

import (
	"fmt"
)

func main() {
	var weight float64 = 75.6
	var weightInt int = int(weight)
	fmt.Printf("Integer: %v, Float: %v\n", weightInt, weight)
}
