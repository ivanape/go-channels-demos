package main

import (
	"fmt"
	"time"
)

func writer(w chan<- int) { // write only channel
	for i := 0; i < 10; i++ {
		w <- i
		time.Sleep(1 * time.Second)
	}

	close(w)
}

func main() {

	channel := make(chan int) // unbuffered channel

	go writer(channel) // start writer goroutine

	// Option 1
	/*for {
		value, ok := <-channel // Blocks execution
		if ok == false { // channel is closed
			fmt.Println("Exit from loop")
			break
		} else {
			fmt.Printf("Value received from channel: %d\n", value)
		}
	}*/

	// Option 2
	for value := range channel {
		fmt.Printf("Value received from channel: %d\n", value)
	}
}
