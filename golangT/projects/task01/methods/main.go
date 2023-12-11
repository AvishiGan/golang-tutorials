package main

import "fmt"

/**********Methods**********/
/**********Basic Method Declaration and Invocation**********/
// 1.
type Rectangle struct {
	length, width float64
}

func (r Rectangle) calculateArea() float64 {
	area := r.length * r.width
	return area
}

// 2.
type Person struct {
	name string
	age  int
}

func (p Person) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}

// 3.
type Book struct {
	title  string
	author string
	year   int
}

func (b Book) displayDetails() {
	fmt.Printf("Title: %s, Author: %s, Year: %d\n", b.title, b.author, b.year)
}

/**********Pointer Receivers and Value Receivers**********/
// 1.
func (r *Rectangle) increaseLengthByTwice() {
	r.length *= 2
}

// 2.
func (p *Person) updateAge(newAge int) {
	p.age = newAge
}

func main() {
	// 1) 1.
	r := Rectangle{length: 5, width: 4}
	fmt.Printf("Area of rectangle: %.2f\n", r.calculateArea())

	// 2) 1.
	r.increaseLengthByTwice()
	fmt.Printf("Updated length of rectangle: %.2f\n", r.length)

	// 1) 2.
	p := Person{name: "Avishi", age: 24}
	p.greet()
	// 2) 2.

	// 1) 3.
	b := Book{title: "Harry Potter", author: "J.K. Rowling", year: 2001}
	b.displayDetails()

}
