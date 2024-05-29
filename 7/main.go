package main

import (
	"fmt"
	"sync"
)

type SyncMap struct {
	data map[int]int
	mu   sync.Mutex
}

// Constructor of concurrent safe map
func NewSyncMap() *SyncMap {
	return &SyncMap{
		data: make(map[int]int),
	}
}

// Adds key value pair to map
func (m *SyncMap) Add(key int, value int) {
	// Locking resource, other goroutines blocks until unlock
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func main() {
	var wg sync.WaitGroup
	m := NewSyncMap()

	wg.Add(3)

	go func() {
		defer wg.Done()
		m.Add(1, 1)
		fmt.Println("goroutine 1 adds to map")
	}()

	go func() {
		defer wg.Done()
		m.Add(2, 2)
		fmt.Println("goroutine 2 adds to map")
	}()

	go func() {
		defer wg.Done()
		m.Add(3, 3)
		fmt.Println("goroutine 3 adds to map")
	}()

	// Waits until all goroutines are done
	wg.Wait()
}
