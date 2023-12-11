package main

import "fmt"

func main() {
	fullName("Avishi", "Ganepola")
}

func fullName(firstName, lastName string) string {
	concatNames := firstName + " " + lastName
	fmt.Println(concatNames)
	return (concatNames)
}
