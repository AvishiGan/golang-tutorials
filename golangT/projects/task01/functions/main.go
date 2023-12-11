package main

import "fmt"

func main() {
	greet()
	add(1, 2)
	calculateCircleArea(5)

	isEven(2)
	calculateVolume(5.2, 3.7, 4.0)
	fullName("Avishi", "Ganepola")

	sum(1, 2, 3, 4, 5, 6, 7, 8, 9)

	// Anonymous function
	// 3) -> 2.
	doubleNumber := func(n int) int {
		return n * 2
	}
	doubledNumber := doubleNumber(4)
	fmt.Printf("Doubled Number: %d\n", doubledNumber)

	concatenate("abcdefghi", "jklmnopqrs", "tuvwxyz")

	// 4) -> 1.

	// 4) -> 2.
	x, y := 10, 20
	swap(&x, &y)

	// 4) -> 3.
	s := []int{1, 2, 3, 4}
	fmt.Println("Slice before modification:", s)
	modifySlice(s)
	fmt.Println("Slice before modification:", s)
}

/**********Basic Function Definition and Invocation**********/

// 1.
func greet() {
	fmt.Println("Hello World!")
}

// 2.
func add(a, b int) int {
	sum := a + b
	fmt.Println(sum)
	return 0
}

// 3.
func calculateCircleArea(radius float64) float64 {
	area := 3.14 * radius * radius
	fmt.Println(area)
	return 0
}

/**********Function Parameters and Return Types**********/

// 1.
func isEven(number int) bool {
	if number%2 == 0 {
		fmt.Printf("\n%d is an even number\n", number)
		return true
	} else {
		fmt.Printf("%d is an odd number\n", number)
		return false
	}
}

// 2.
func calculateVolume(length, width, height float64) float64 {
	volume := length * width * height
	fmt.Printf("Volume = %f\n", volume)
	return 0
}

// 3.
func fullName(firstName, lastName string) string {
	concatNames := firstName + " " + lastName
	fmt.Println(concatNames)
	return (concatNames)
}

/**********Variadic Functions and Anonymous Functions**********/

// 1.
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Printf("\nTotal = %d\n", total)
	return total
}

// 3.
func concatenate(strings ...string) string {
	var result string

	for _, str := range strings {
		result += str
	}
	fmt.Println(result)
	return (result)
}

/**********Function Parameters - Passing by Value and Reference**********/
// 1.
func modifySlice(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

// 2.
func swap(x, y *int) {
	fmt.Printf("Before swap x: %d, y: %d\n", *x, *y)
	*x, *y = *y, *x
	fmt.Printf("After swap x: %d, y: %d\n", *x, *y)
}
