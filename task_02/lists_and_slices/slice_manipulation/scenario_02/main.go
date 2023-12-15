package main

import "fmt"

func main() {
	names := []string{"Alice", "Bob", "Charlie", "Pete", "Dave", "Kayla"}
	fmt.Printf("Names slice before: %v\n", names)

	names = append(names[:2], names[3:]...)
	fmt.Printf("Names slice after removing the 2nd element: %v", names)
}
