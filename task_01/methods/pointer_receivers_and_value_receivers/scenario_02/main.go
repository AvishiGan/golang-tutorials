package main

import "fmt"

type Book struct {
	title  string
	author string
	year   int
}

func (b *Book) changeAuthor(newAuthor string) {
	b.author = newAuthor
}

func main() {
	b := Book{title: "Harry Potter", author: "J.K. Rowling", year: 2001}
	b.changeAuthor("Emma Watson")
	fmt.Printf("New author: %s\n", b.author)
}
