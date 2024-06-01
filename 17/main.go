package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	arr := []int{1, 5, 13, 24, 67, 83, 103}
	t, err := binarySearch(arr, 24)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("found at index:", t)
}

func binarySearch(arr []int, target int) (int, error) {
	// Initialize low and high pointers
	low := 0
	high := len(arr) - 1

	// Continue searching while low pointer is less than or equal to high pointer
	for low <= high {
		// Calculate mid index
		mid := low + (high-low)/2

		// If target found at mid index, return mid
		if arr[mid] == target {
			return mid, nil
		} else if arr[mid] < target { // If target is greater, search in the right half
			low = mid + 1
		} else { // If target is smaller, search in the left half
			high = mid - 1
		}
	}

	return -1, errors.New("target not found")
}
