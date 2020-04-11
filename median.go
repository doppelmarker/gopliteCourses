package main

import (
	"fmt"
	"sort"
)

func median(array []int) float64 {
	if len(array)%2 != 0 {
		return float64(array[len(array)/2])
	}
	return float64((array[len(array)/2] + array[len(array)/2-1])) / 2
}

func main() {
	var array1 []int = []int{6, 4, 2, 3, 5, 1} //Even amount of numbers
	sort.Ints(array1)
	fmt.Println(median(array1))
	var array2 []int = []int{5, 0, 1, 4, 16} //Odd amount of numbers
	sort.Ints(array2)
	fmt.Print(median(array2))
}
