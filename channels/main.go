package main

import (
	"fmt"
	"time"
)

// channels are thge way goroutines communicate with each other.
func main() {
	// In count program we ar directly getting output in terminal, but what if we want to communicate back to the main go routing
	// we accept channel as an argument, and channel is like a pipe through which you can send and recieve a message, channels have any "types"

	// goroutines helps the program to run concurrently
	// channels allows us to synchronize goroutines
	c := make(chan string)
	// recieving
	go Count("hello", c)
	// for {
	// 	msg, open := <-c
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }
	// or
	for msg := range c {
		fmt.Println(msg)
	}

}

// sending
func Count(things string, c chan string) {
	for i := 0; i <= 5; i++ {
		c <- things
		//fmt.Println(things, i)
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
