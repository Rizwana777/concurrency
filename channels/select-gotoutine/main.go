package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "first 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "second 2 seconds"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	// using just for loop will take time so we have to use select statement
	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
