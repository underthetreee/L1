package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.RWMutex
}

// Increment counter
func (c *Counter) Increment() {
	// Locks mutex, provides exclusive write access to resource
	// others goroutines waits until unlock
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Return counter value
func (c *Counter) Value() int {
	// Goroutines which rlocks, can read concurrently safe
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func main() {
	var (
		wg sync.WaitGroup
		c  Counter
	)

	// Launch 100 goroutines to increment counter concurrently
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()

	}
	// Waits until all goroutines are done
	wg.Wait()
	fmt.Println("counter:", c.Value())
}
