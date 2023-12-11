package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*************Variables and Constants**************/
	fmt.Println("Variables and Constants")

	/*************Variables**************/
	fmt.Println("Variables")

	// 1.
	var age int = 25
	fmt.Printf("1.	Value: %v, Type: %T\n", age, age)

	// 2.
	var name string = "Alice"
	var city string = "New York"
	fmt.Println("2.	" + name + " lives in " + city)

	// 3.
	var temperature float64 = 78.5
	var tempInt int = int(temperature)
	fmt.Printf("3.	Float: %v, Int: %v\n", temperature, tempInt)

	// 4.
	var isRunning bool = true
	if isRunning == true {
		println("4.	Take an umbrella")
	}

	// 5.
	fmt.Println("5")
	var a int = 10
	var b int = 20
	fmt.Printf("Before swap a:%v, b:%v\n", a, b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("After swap a:%v, b:%v\n\n", a, b)

	/*************Constants**************/
	fmt.Println("Constants")
	// 1.
	const pi = 3.14159
	fmt.Printf("1.	pi: %v\n", pi)

	// 2.
	const (
		Monday    = "Monday"
		Tuesday   = "Tuesday"
		Wednesday = "Wednesday"
		Thursday  = "Thursday"
		Friday    = "Friday"
		Saturday  = "Saturday"
		Sunday    = "Sunday"
	)
	fmt.Println("2.")
	fmt.Printf(Monday)
	fmt.Printf("\n")
	fmt.Printf(Tuesday)
	fmt.Printf("\n")
	fmt.Printf(Wednesday)
	fmt.Printf("\n")
	fmt.Printf(Thursday)
	fmt.Printf("\n")
	fmt.Printf(Friday)
	fmt.Printf("\n")
	fmt.Printf(Saturday)
	fmt.Printf("\n")
	fmt.Printf(Sunday)
	fmt.Printf("\n")

	// 3.
	const hoursInDay = 24
	var numberOfMinutesPerDay int = hoursInDay * 60
	fmt.Printf("Number of minutes in a day = %v\n", numberOfMinutesPerDay)

	// 4.
	const taxRate = 0.08
	var purchase int = 120
	fmt.Printf("Tax for %v purchase = %v\n", purchase, taxRate*120)

	// 5.
	const version = "v1.2.3"
	fmt.Printf("5.	Current version is %v\n\n", version)

	/*************Conversions**************/
	fmt.Println("Conversions")

	/*************Type Conversions**************/
	fmt.Println("Type Conversions")

	// 1.
	var num1 int = 42
	var num1Float float64 = float64(num1)
	fmt.Printf("1.	int:%v, float:%v\n", num1, num1Float)

	// 2.
	var num string = "1234"
	numInt, err := strconv.Atoi(num) // Atoi stands for ASCII to Integer
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("2.	The integer is %v\n", numInt)
	}

	// 3.
	var weight float64 = 75.6
	var weightInt int = int(weight)
	fmt.Printf("3.	Integer: %v, Float: %v\n", weightInt, weight)

	// 4.
	var status bool = true
	var statusInt int
	if status {
		statusInt = 1
	} else {
		statusInt = 0
	}
	fmt.Printf("3.	Boolean: %v, Int: %v\n\n", status, statusInt)

	/*************String Conversions**************/
	fmt.Println("String Conversions")

	// 1.
	var count int = 16
	var countString string = strconv.Itoa(count) // Itoa -> Converts an integer to a string in base 10
	var str_var string = "cats"
	fmt.Printf("1.	" + countString + " " + str_var)

	// 2.
	fmt.Printf("\n2.	")
	s := "3.14" // Short variable declaration
	var f float64

	f, err = strconv.ParseFloat(s, 10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}

	// 3.
	fmt.Printf("3.	")
	distance := 40.10
	distanceString := fmt.Sprintf("%f", distance)
	fmt.Println(distanceString)

	// 4.
	fmt.Printf("4.	")
	isValid := true
	isValidString := fmt.Sprintf("%t", isValid)
	str_concat := "That is "
	fmt.Printf(str_concat + isValidString)

	// 5.
	fmt.Printf("\n5.	")
	year := 2023
	yearString := fmt.Sprintf("%d", year)
	fmt.Printf("The current year is %s", yearString)

	// 6.
	fmt.Printf("\n6.	")
	stringInput := "1234"

	stringInputInt, err := strconv.Atoi(stringInput)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(stringInputInt)
	}
}
