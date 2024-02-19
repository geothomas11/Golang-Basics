package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mutex sync.Mutex

func main() {
	// Number of goroutines to create
	numGoroutines := 3

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Create multiple goroutines
	for i := 1; i <= numGoroutines; i++ {
		go incrementCounter(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Final Counter Value:", counter)
}

func incrementCounter(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		// Lock the mutex before accessing the shared counter
		mutex.Lock()

		// Simulate some work
		time.Sleep(time.Millisecond * 100)

		// Increment the counter
		counter++

		// Unlock the mutex after updating the counter
		mutex.Unlock()

		// Simulate some additional work
		time.Sleep(time.Millisecond * 50)
		fmt.Printf("Goroutine %d: Counter Value %d\n", id, counter)
	}
}
