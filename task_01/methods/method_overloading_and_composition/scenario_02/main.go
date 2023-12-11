package main

import "fmt"

type Car struct {
	name string
}

func (c Car) drive() {
	fmt.Printf("%s is being driven", c.name)
}

type ElectricCar struct {
	Car
}

func main() {
	eCar := ElectricCar{
		Car: Car{
			name: "Tesla Model S",
		},
	}
	eCar.drive()
}
