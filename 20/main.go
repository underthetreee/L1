package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "snow dog sun"
	reversed := reverse(s)
	fmt.Println(reversed)
}

func reverse(s string) string {
	// Slicing string
	words := strings.Fields(s)
	reversed := make([]string, len(words))
	for i, w := range words {
		// Write words in reversed order
		reversed[len(words)-1-i] = w
	}
	// Joining words into string
	return strings.Join(reversed, " ")
}
