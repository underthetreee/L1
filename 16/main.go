package main

import "fmt"

func main() {
	arr := []int{3, 6, 11, 24, 1, 19}
	fmt.Println("Unsorted array:", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		// Partitioning index
		pi := partition(arr, low, high)

		// Recursively sort elements before and after partition
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// Pivot (Element to be placed at right position)
	pivot := arr[high]

	i := low - 1 // Index of smaller element

	for j := low; j < high; j++ {
		// If current element is smaller than the pivot
		if arr[j] < pivot {
			i++
			// Swap arr[i] and arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// Swap arr[i+1] and arr[high]
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
