package main

import (
	"fmt"
	"time"
)

/*
	Main calling two go routines:
		- listener: shows numbers sent to channel
		- writer: writes numbers into channel
*/

func listener(l <-chan int) { // read only channel
	for {
		value := <-l
		fmt.Printf("Value received: %d\n", value)
	}
}

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

	go listener(channel) // start listener goroutine
	go writer(channel)   // start writer goroutine

	// Block execution
	fmt.Println("Press any key to exit...")
	var exit string
	_, _ = fmt.Scanln(&exit)

	close(channel) // close channel
}
