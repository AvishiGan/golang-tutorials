package main

import "fmt"

type Speaker interface {
	speak()
}

type Person struct {
	name string
}

func (p Person) speak() {
	fmt.Printf("%s can speak more than 3 languages", p.name)
}

func main() {
	p := Person{name: "Avishi"}
	p.speak()
}
