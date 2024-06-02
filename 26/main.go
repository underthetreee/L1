package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aabcd"
	fmt.Println(IsUnique(s))
}

func IsUnique(str string) bool {
	// Convert string to lowercase
	str = strings.ToLower(str)
	// Map to track encountered characters
	charMap := make(map[rune]bool)

	for _, char := range str {
		// If encounted before return false
		if charMap[char] {
			return false
		}
		// Mark character as encounted
		charMap[char] = true
	}
	// If all characters are unique return true
	return true
}
