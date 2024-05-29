package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	result := make(chan int)

	// Iterates over a slice of numbers, calling a goroutine on each iteration
	// to send square of number to result channel
	for _, n := range numbers {
		wg.Add(1)
		n := n
		go func() {
			defer wg.Done()
			result <- n * n
		}()
	}

	// Waits until goroutines are done, then closes result channel
	go func() {
		wg.Wait()
		close(result)
	}()

	// Reads from result channel, then adds to sum variable
	sum := 0
	for r := range result {
		sum += r
	}
	fmt.Println(sum)
}
