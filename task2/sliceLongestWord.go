package main

import "fmt"

//max is used to find the longest word from "slice" (the first if there are more than one)
func max(slice []string) string {
	max := slice[0]
	for _, elem := range slice {
		if len(elem) > len(max) {
			max = elem
		}
	}
	return max
}

func sliceLongestWord() {
	slice1 := []string{"Du", "Riechst", "So", "Gut"}
	fmt.Println("Slice1 longest word:", max(slice1))
	slice2 := []string{"oh", "hi", "M"}
	fmt.Println("Slice2 longest word:", max(slice2))
}
