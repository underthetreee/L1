package main

import "fmt"

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}

	set := strSet(s)

	// Iterates over set and writing values to stdout
	for _, s := range set {
		fmt.Println(s)
	}
}

// Takes slice of strings and returns slice with unique strings
func strSet(strs []string) []string {
	set := make(map[string]bool)
	var result []string

	// Iterates over slice with strings
	for _, s := range strs {
		// If current string is not already in map,
		// then adds to map	as key with true value
		// and append to result slice
		if !set[s] {
			set[s] = true
			result = append(result, s)
		}
	}
	return result
}
