package main

import (
	"fmt"
	"time"
)

func writer(w chan<- int) { // writer only channel
	time.Sleep(3 * time.Second)
	w <- 3
	time.Sleep(3 * time.Second)
	w <- 4
}

func main() {

	channel := make(chan int) // unbuffered channel

	go writer(channel) // start writer goroutine

	// This assigment is blocking
	value1 := <-channel
	fmt.Printf("Value received from channel: %d\n", value1)
	value2 := <-channel
	fmt.Printf("Value received from channel: %d\n", value2)
	close(channel)
}
