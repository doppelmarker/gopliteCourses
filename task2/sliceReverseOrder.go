package main

import "fmt"

//reverse returns reversed copy of "slice"
func reverse(slice []int64) []int64 {
	c := make([]int64, len(slice))
	copy(c, slice)
	for i, j := 0, len(c)-1; i < j; i++ {
		c[i], c[j] = c[j], c[i]
		j--
	}
	return c
}

func sliceReverseOrder() {
	slice := []int64{6, 0, 1, 3, 5, 4, 8}
	fmt.Println("The copy of the original slice in reverse order: ", reverse(slice))
	fmt.Println("Original slice:", slice)
}
