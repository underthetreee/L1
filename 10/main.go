package main

import (
	"fmt"
	"slices"
)

const step = 10.0

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempGroups := make(map[int][]float64)

	// Sorts slice
	slices.Sort(temps)

	// Iterates over temps, calculating key based on step size
	// adds temps to related group
	for _, t := range temps {
		key := int(t/step) * 10
		tempGroups[key] = append(tempGroups[key], t)
	}

	// Iterates over map and writing k, v to stdout
	for k, v := range tempGroups {
		fmt.Printf("%d: %v\n", k, v)
	}
}
