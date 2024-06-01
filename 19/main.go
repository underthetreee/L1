package main

import "fmt"

func main() {
	str := "главрыба"
	r := reverseString(str)
	fmt.Println(str, r)
}

// Takes string and returns its reversed version
func reverseString(str string) string {
	// Convert string to slice of runes
	r := []rune(str)
	// Use i and j to traverse array from both ends
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		// Swap elements i and j
		r[i], r[j] = r[j], r[i]
	}
	// Convert slice of rune to string and return it
	return string(r)
}
