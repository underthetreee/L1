package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	duration := 1 * time.Second
	dataCh := make(chan int)

	sendData(dataCh, duration)

	// Reads data from channel, then writing it to stdout
	for data := range dataCh {
		fmt.Println("received:", data)
	}
}

// Sends data to channel until time is out
func sendData(dataCh chan<- int, duration time.Duration) {
	go func() {
		defer close(dataCh)

		timer := time.NewTimer(duration)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				fmt.Println("time is expired")
				return
			case dataCh <- genData():
			}
		}
	}()

}

// Generate random integer from 0 to 100
func genData() int {
	return rand.Intn(100)
}
