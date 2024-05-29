package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Flag to set number of workers
	workers := flag.Int("w", 3, "number of workers")
	flag.Parse()

	var (
		dataCh      = make(chan int)
		interruptCh = make(chan os.Signal)
		doneCh      = make(chan struct{})
	)

	signal.Notify(interruptCh, syscall.SIGINT)

	// Waits interrupt signal to exit the program
	go func() {
		<-interruptCh
		close(doneCh)
	}()

	// Adds workers to wait group
	wg := &sync.WaitGroup{}
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		receiveData(wg, dataCh, doneCh)
	}

	// Sends random integers to channel
	go func() {
		defer close(dataCh)
		for {
			dataCh <- genData()
		}
	}()

	// Waits until all workers are done
	wg.Wait()
}

// Generate random integer from 0 to 100
func genData() int {
	return rand.Intn(100)
}

// Reads data from channel and writing it to stdout
func receiveData(wg *sync.WaitGroup, dataCh <-chan int, doneCh <-chan struct{}) {
	go func() {
		defer wg.Done()
		for {
			select {
			case <-doneCh:
				return
			case data := <-dataCh:
				fmt.Println("received", data)
			}
		}
	}()
}
