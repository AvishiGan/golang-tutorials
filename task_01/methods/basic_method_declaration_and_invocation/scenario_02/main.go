package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}

func main() {
	p := Person{name: "Avishi", age: 24}
	p.greet()
}
