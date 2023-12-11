package main

import "fmt"

// Structs -> data, interfaces -> abstract behavior
type Vehicle interface {
	startEngine()
	stopEngine()
}

type Car struct {
	name string
}

func (c Car) startEngine() {
	fmt.Printf("%s has started the engine\n", c.name)
}

func (c Car) stopEngine() {
	fmt.Printf("%s has stopped the engine\n", c.name)
}

func main() {
	c := Car{name: "BMW i3"}
	c.startEngine()
	c.stopEngine()
}
