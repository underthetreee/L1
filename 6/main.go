package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Stops goroutine with empty struct channel
	doneCh := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-doneCh:
			fmt.Println("stop goroutine")
			return
		}
	}()
	// Or doneCh<-struct{}{}
	close(doneCh)

	// Stops goroutine with context cancel
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("stop goroutine 2")
			return
		}
	}()
	cancel()

	// Stops goroutine with timer
	wg.Add(1)
	go func() {
		defer wg.Done()
		timer := time.NewTimer(2 * time.Second)
		defer timer.Stop()
		select {
		case <-timer.C:
			fmt.Println("stop goroutine 3")
			return
		}
	}()

	// Waits until all goroutines are done
	wg.Wait()
}
