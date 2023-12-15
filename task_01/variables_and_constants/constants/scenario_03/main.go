package main

import (
	"fmt"
)

func main() {
	const hoursInDay = 24
	var numberOfMinutesPerDay int = hoursInDay * 60
	fmt.Printf("Number of minutes in a day = %v\n", numberOfMinutesPerDay)
}
