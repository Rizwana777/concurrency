package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutines helps the program to run concurrently
	go Count("hello")
	go Count("world")

	fmt.Scanln()
}

func Count(things string) {
	for i := 0; true; i++ {
		fmt.Println(things, i)
		time.Sleep(time.Millisecond * 500)
	}
}
