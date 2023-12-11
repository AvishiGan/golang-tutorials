package main

import (
	"fmt"
)

func main() {
	const taxRate = 0.08
	var purchase int = 120
	fmt.Printf("Tax for %v purchase = %v\n", purchase, taxRate*120)
}
