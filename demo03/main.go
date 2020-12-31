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

	channel1 := make(chan int) // unbuffered channel1
	channel2 := make(chan int) // unbuffered channel1

	go writer(channel1) // start writer goroutine
	go writer(channel2) // start writer goroutine

	for i := 0; i < 10; i++ {
		select { // Blocks execution until we receive information from any of the channels
		case value := <-channel1:
			fmt.Printf("Value received from channel1: %d\n", value)
		case value := <-channel2:
			fmt.Printf("Value received from channel2: %d\n", value)
		}
	}

	close(channel1) // close channel1
	close(channel2) // close channel2
}
