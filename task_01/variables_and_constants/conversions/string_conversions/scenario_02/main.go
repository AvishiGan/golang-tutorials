package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "3.14" // Short variable declaration
	var f float64
	var err error

	f, err = strconv.ParseFloat(s, 10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}
}
