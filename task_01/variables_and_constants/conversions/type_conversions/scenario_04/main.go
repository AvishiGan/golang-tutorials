package main

import (
	"fmt"
)

func main() {
	var status bool = true
	var statusInt int
	if status {
		statusInt = 1
	} else {
		statusInt = 0
	}
	fmt.Printf("Boolean: %v, Int: %v\n", status, statusInt)
}
