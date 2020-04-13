package main

import (
	"fmt"
	"sort"
)

//printSorted prints "m" values sorted in order of increasing keys
func printSorted(m map[int]string) {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Print(m[k], " ")
	}
}

func mapSortedKeys() {
	testMap := map[int]string{3: "das", 1: "Wollt", 4: "Bett", 2: "ihr", 5: "in", 7: "sehen", 6: "Flammen"}
	fmt.Println("Map's values sorted in order of increasing keys: ")
	printSorted(testMap)
}
