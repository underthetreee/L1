package main

import "fmt"

func main() {
	var (
		a = 5
		b = 10
	)

	a, b = arithmeticSwap(a, b)
	fmt.Println(a, b)

	a, b = bitwiseSwap(a, b)
	fmt.Println(a, b)
}

// Swaps value of a and b using arithmetic operations
func arithmeticSwap(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

// Swaps value of a and b using XOR operations
func bitwiseSwap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}
