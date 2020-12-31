package main

import (
	"fmt"
	"sync"
)

func listener(wg *sync.WaitGroup, r <-chan int) { // read only channel
	defer wg.Done()
	for i := 0; i < 3; i++ {
		value := <-r
		fmt.Printf("Value received from channel: %d\n", value)
	}
}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)
	channel := make(chan int, 3) // buffered channel

	go listener(&wg, channel) // start listener goroutine

	// This assigment is not blocking, so we need to use a WaitGroup
	channel <- 1
	channel <- 2
	channel <- 3

	wg.Wait()
	fmt.Println("Exit main")
	close(channel)
}
