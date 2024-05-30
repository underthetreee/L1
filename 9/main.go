package main

import "fmt"

func main() {
	nums := []int{2, 4, 8, 16, 32, 64}

	var (
		in  = make(chan int)
		out = make(chan int)
	)

	// Iterates over slice and sending values into channel
	go func() {
		for _, n := range nums {
			in <- n
		}
		close(in)
	}()

	double(in, out)

	// Reads from channel and prints values to stdout
	for n := range out {
		fmt.Println(n)
	}
}

// Reads values from "in" channel, then doubles it and sends it to "out" channel
func double(in <-chan int, out chan<- int) {
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
}
