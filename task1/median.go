package main

import (
	"fmt"
	"sort"
)

func median(array []int) float64 {
	sort.Ints(array)
	if len(array)%2 != 0 {
		return float64(array[len(array)/2])
	}
	return float64((array[len(array)/2] + array[len(array)/2-1])) / 2
}

func medianTask() {
	var array1 []int = []int{6, 4, 2, 3, 5, 1}
	fmt.Println("Median for even number of elements: ", median(array1))
	var array2 []int = []int{5, 0, 1, 4, 16}
	fmt.Println("Median for odd number of elements: ", median(array2))
}
