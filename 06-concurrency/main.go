package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// in Go, we can use the `go` keyword to run a function
	// at a different "thread". This is called "launching a goroutine".
	// It's not quite a thread but works similarly for our purposes.
	// When launching a goroutine, make sure to have a way of waiting for
	// it to complete, otherwise your program will exit without it fully executing.
	// The most basic way of doing that is using the `sync` package.
	wg := &sync.WaitGroup{}

	// When using a WaitGroup, you need to inform it how many parallel works you're running
	// Then calling Done for each one that completes.
	wg.Add(1)
	go countToTen(wg, 1)

	wg.Add(1)
	go countToTen(wg, 2)

	// Wait is going to block the main function until wg.Done() is called twice.
	wg.Wait()

	// A more efficient way of doing that is passing in a channel of strings
	// and listening on it to print the messages.
	// Channels work like a pub/sub mechanism, as you can send messages to it
	// and someone else can read from them. Channels block when full for writes,
	// and when empty for reads. You can use many concurrent patterns like
	// FanIn, FanOut and Pipelines by using channels with goroutines.
	result1 := countToTenWithChannel(1)
	result2 := countToTenWithChannel(2)

	result := fanIn(result1, result2)

	for {
		msg := <-result
		if msg == "" {
			break
		}

		fmt.Println(msg)
	}
}

// this function will count to 10, waiting 1s between iterations
func countToTen(wg *sync.WaitGroup, processID int) {
	// The keyword `defer` puts the function passed to it at the end of the
	// `stack` for the current function executing. This makes sure that the
	// wg.Done() is going to be executed before exiting the countToTen function.
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(fmt.Sprintf("process %d counting %d", processID, i))
	}
}

// this function will count to 10, waiting 1s between iterations
// and pushing the result to a channel
func countToTenWithChannel(processID int) <-chan string {
	result := make(chan string)
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			result <- fmt.Sprintf("process %d counting %d and sending to channel", processID, i)
		}
		// closing a channel signals to consumers that they shouldn't
		// block on reads, because there's not going to be any more messages.
		close(result)
	}()

	return result
}

// this function reads from both channels, merging the results
func fanIn(result1, result2 <-chan string) <-chan string {
	result := make(chan string)
	go func() {
		for {
			select {
			case msg := <-result1:
				result <- msg
			case msg := <-result2:
				result <- msg
			}
		}
	}()

	return result
}
