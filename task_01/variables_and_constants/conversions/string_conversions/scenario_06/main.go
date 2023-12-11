package main

import (
	"fmt"
	"strconv"
)

func main() {
	stringInput := "1234"

	stringInputInt, err := strconv.Atoi(stringInput)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(stringInputInt)
	}
}
