package main

import (
	"strings"
)

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }
//
// func main() {
//   someFunc()
// }

var justString string

// Create string with repeated symbols of provided length
func createHugeString(length int) string {
	s := strings.Repeat("s", length)
	return s
}

func someFunc() {
	// Return string of 1024 bytes
	v := createHugeString(1 << 10)
	// In source example if length less of 100 will cause out of range
	if len(v) > 100 {
		// Create slice of 100 bytes
		b := make([]byte, 100)
		// Copy 100 bytes of v into b slice
		// In source example justString will points to original v string,
		// GC will not clean up it, that why needs to make copy
		copy(b, v[:100])
		justString = string(b)
	}
}

func main() {
	someFunc()
}
