package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Println("Slice before modification:", s)
	modifySlice(s)
	fmt.Println("Slice before modification:", s)
}

func modifySlice(s []int) {
	for i := range s {
		s[i] *= 2
	}
}
