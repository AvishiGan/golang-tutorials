package main

import (
	"fmt"
)

func main() {
	var temperature float64 = 78.5
	var tempInt int = int(temperature)
	fmt.Printf("Float: %v, Int: %v\n", temperature, tempInt)
}
