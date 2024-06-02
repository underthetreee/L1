package main

import "time"

func main() {
	sleep(5)
}

func sleep(seconds int) {
	// Return channel after duration
	// Blocks until duration elapses
	<-time.After(time.Duration(seconds) * time.Second)
}
