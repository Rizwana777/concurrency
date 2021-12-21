package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for.
	// Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.
	// A WaitGroup must not be copied after first use.

	var wg sync.WaitGroup
	// Add adds delta, which may be negative, to the WaitGroup counter. If the counter becomes zero, all goroutines blocked on Wait are released.
	// If the counter goes negative, Add panics.Note that calls with a positive delta that occur when the counter is zero must happen before a Wait.
	// Calls with a negative delta, or calls with a positive delta that start when the counter is greater than zero, may happen at any time.
	// Typically this means the calls to Add should execute before the statement creating the goroutine or other event to be waited for.
	// If a WaitGroup is reused to wait for several independent sets of events, new Add calls must happen after all previous Wait calls have returned.
	// See the WaitGroup example.
	wg.Add(1)
	go func() {
		Count("sheep")
		Count("hello")
		wg.Done()
	}()
	go func() {
		Count("hello")
		wg.Done()
	}()
	wg.Wait()
}

func Count(things string) {
	for i := 0; true; i++ {
		fmt.Println(things, i)
		time.Sleep(time.Millisecond * 500)
	}
}
