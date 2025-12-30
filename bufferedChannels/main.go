package main

import "fmt"

func main() {
	fmt.Println("Buffered Channels")

	bfChannel := make(chan int, 2)

	bfChannel <- 1
	bfChannel <- 2

	fmt.Println(<-bfChannel)
	bfChannel <- 3
	fmt.Println(<-bfChannel)
	bfChannel <- 4
	fmt.Println(<-bfChannel)
	fmt.Println(<-bfChannel)

	bfChannelsWithLoops()

	fmt.Println("Bye")
}

func bfChannelsWithLoops() {
	bfChannel := make(chan int, 2)

	go func() {
		for i := 0; i < 10; i++ {
			bfChannel <- i
			bfChannel <- i + 1
		}
		close(bfChannel)
	}()

	for {
		message, isOpen := <-bfChannel
		if !isOpen {
			break
		}
		fmt.Println(message)
	}
}
