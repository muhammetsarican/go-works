package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Select")

	ch := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch <- "Channel 1"
	}()

	go func() {
		ch2 <- "Channel 2"
	}()

	var msg1 string
	var msg2 string

	for len(msg1) == 0 || len(msg2) == 0 {
		select {
		case msg1 = <-ch:
			fmt.Println("Received", msg1)
		case msg2 = <-ch2:
			fmt.Println("Received", msg2)
		default:
			fmt.Println("No message received")
		}
		time.Sleep(1 * time.Second)
	}
}
