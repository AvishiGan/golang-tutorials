package main

import "fmt"

func main() {
	calculateVolume(5.2, 3.7, 4.0)
}

func calculateVolume(length, width, height float64) float64 {
	volume := length * width * height
	fmt.Printf("Volume = %f\n", volume)
	return 0
}
