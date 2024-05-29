package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	result := make(chan int)

	for _, n := range numbers {
		wg.Add(1)
		n := n
		go func() {
			defer wg.Done()
			result <- n * n
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	sum := 0
	for r := range result {
		sum += r
	}
	fmt.Println(sum)
}
