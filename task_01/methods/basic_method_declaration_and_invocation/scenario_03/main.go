package main

import "fmt"

type Book struct {
	title  string
	author string
	year   int
}

func (b Book) displayDetails() {
	fmt.Printf("Title: %s, Author: %s, Year: %d\n", b.title, b.author, b.year)
}

func main() {
	b := Book{title: "Harry Potter", author: "J.K. Rowling", year: 2001}
	b.displayDetails()
}
