package main

import "fmt"

//average is used to find an average value of "array"
func average(array [6]int) float64 {
	var average float64
	for _, elem := range array {
		average += float64(elem)
	}
	average /= 6
	return average
}

func arrayAverage() {
	array := [6]int{3, 4, 5, 6, 7, 1}
	fmt.Println("Array's average value:", average(array))
}
