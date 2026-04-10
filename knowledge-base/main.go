package main

import (
	"fmt"
	"sort"
)

func main() {
	// Slices and variadic functions:
	reminder()
	addUsers("Alice", "Bob", "Charlie")
	err := removeUsers(1)
	if err != nil {
		fmt.Println("Error:", err)
	}
	numbers := Myslice{1, 2, 3, 4, 5}
	numbers, err = numbers.Remove(2)
	if err != nil {
		fmt.Println("Error:", err)
	}

	toSortNumber := SortNumbers{5, 2, 9, 1, 5, 6}
	fmt.Println("Unsorted numbers:", toSortNumber)
	// sort.Sort(toSortNumber)
	sort.Sort(byInc{toSortNumber})
	fmt.Println("Sorted numbers:", toSortNumber)

	toSortNumber2 := SortNumbers{5, 2, 9, 1, 5, 6}
	sort.Sort(byDec{toSortNumber2})
	fmt.Println("Sorted numbers (descending):", toSortNumber2)
	// Maps functions:
}
