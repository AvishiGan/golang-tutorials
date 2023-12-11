package main

import "fmt"

type Calculator struct {
	result int
}

func (c *Calculator) add(value int) *Calculator {
	c.result += value
	return c
}

func (c *Calculator) subtract(value int) *Calculator {
	c.result -= value
	return c
}

func (c *Calculator) multiply(value int) *Calculator {
	c.result *= value
	return c
}

func main() {
	calc := Calculator{}
	calc.add(100).subtract(55).multiply(4)
	fmt.Printf("Result is %v", calc.result)
}
