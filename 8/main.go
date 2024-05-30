package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64 = 60

	result := setBit(2, n)
	fmt.Println(strconv.FormatInt(n, 2))
	fmt.Println(strconv.FormatInt(result, 2))
}

// Shifts bit to specified position, then invert bit
func setBit(pos uint, n int64) int64 {
	return n ^ (1 << pos)
}
