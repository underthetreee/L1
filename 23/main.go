package main

import "fmt"

func main() {
	s := []int{15, 23, 55, 11}
	fmt.Println("source slice:", s)

	i := 1
	// Append with slice before deleting element and slice after deleting element
	s = append(s[:i], s[i+1:]...)
	fmt.Println("after delete:", s)
}
