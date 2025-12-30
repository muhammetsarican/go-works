package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan string)

	chDone := make(chan bool)

	go func() {
		ch <- 1
		chDone <- true
	}()
	go func() {
		ch2 <- "waiting..."
		chDone <- true
	}()
	go func() {
		ch3 <- "still waiting"
		chDone <- true
	}()

	message, isOpen := <-ch2
	fmt.Println(message, isOpen)
	fmt.Println(<-ch)
	fmt.Println(<-ch3)

	<-chDone
	<-chDone
	<-chDone

	fmt.Println("Bye")

	unbufferedChannelsWithLoop()
}

func unbufferedChannelsWithLoop() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		message, isOpen := <-ch
		if !isOpen {
			break
		}
		fmt.Println(message)
	}

	fmt.Println("Bye")
}
