package main

import (
	"fmt"
	"time"
)

func writer(w chan<- int) { // write only channel
	value := 0
	for {
		w <- value
		time.Sleep(1 * time.Second)
		value++
	}
}

func main() {

	channel := make(chan int) // unbuffered channel

	go writer(channel) // start writer goroutine

	for i := 0; i < 10; i++ {
		value := <-channel 	// Blocks execution
		fmt.Printf("Value received: %d\n", value)
	}

	close(channel) // close channel
}
