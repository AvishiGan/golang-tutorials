package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b int = 20
	fmt.Printf("Before swap a:%v, b:%v\n", a, b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("After swap a:%v, b:%v\n\n", a, b)
}
