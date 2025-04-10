// 09: Conurrency: Goroutines and Channels

package main

import (
	"fmt"
	"sync"
	"time"
)

// Goroutines and Channels:
// - A goroutine is a lightweight thread managed by the Go runtime.
// - It allows functions to run concurrently with other functions.
// - A channel is a Go feature that allows goroutines to communicate and synchronize by passing data.

func sayHello() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Hello", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func sendMessage(ch chan string) {
	ch <- "Hello from Goroutine!"
}

// Worker function that simulates doing some work
func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting...\n", id)
	time.Sleep(time.Second) // simulate some work
	fmt.Printf("Worker %d done.\n", id)

	wg.Done() // Notify that this worker is done
}

func main() {
	// Goroutines:
	go sayHello() // Runs concurrently
	fmt.Println("Main function")
	time.Sleep(3 * time.Second) // Give time for goroutine to finish

	// Channels:
	ch := make(chan string) // Create a channel
	go sendMessage(ch)      // Start goroutine
	msg := <-ch             // Receive message
	fmt.Println(msg)

	// WaitGroup: Use sync.WaitGroup to wait for multiple goroutines to finish.
	var wg sync.WaitGroup

	// Launch 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Tell WaitGroup to wait for 1 more goroutine
		go worker(i, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers completed.")
}

/**
Output of Goroutines:
- Main function
- Hello 1
- Hello 2
...

Output of Channels:
- Hello from Goroutine!

Output of WaitGroup:
- Worker 1 starting...
- Worker 2 starting...
- Worker 3 starting...
- Worker 3 done.
- Worker 2 done.
- Worker 1 done.
- All workers completed.

# How It Works:
- wg.Add(1) is called before each go worker(...).
- Each worker() finishes and calls wg.Done() to decrement the count.
- wg.Wait() blocks the main() function until all .Done() calls match the .Add() calls.
The order may vary slightly because goroutines run concurrently.
*/
