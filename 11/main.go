package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 14, 29}
	nums2 := []int{7, 5, 2, 3, 1}

	res := intersection(nums, nums2)
	fmt.Println(res)
}

func intersection(first []int, second []int) []int {
	var result []int
	setMap := make(map[int]bool)

	// Iterates over first slice and set elements as keys with true value
	for _, n := range first {
		setMap[n] = true
	}

	// Iterates over second slice checking if element exists in map
	// appends to result slice if it exists
	for _, n := range second {
		if setMap[n] {
			result = append(result, n)
		}
	}
	return result
}
