package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Go Routines and WaitGroups")

	wg := sync.WaitGroup{}

	wg.Add(3)
	go say("hello")
	wg.Done()

	go say("hello")
	wg.Done()

	go func() {
		defer wg.Done()
		fmt.Println("defer example")
	}()

	wg.Wait()

	fmt.Println("Bye")

	goRoutineExample()
}

func say(word string) {
	fmt.Println(word)
}

func goRoutineExample() {
	startTime := time.Now()

	wg := sync.WaitGroup{}

	wg.Add(3)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println(time.Now().Sub(startTime))
	}()
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println(time.Now().Sub(startTime))
	}()
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println(time.Now().Sub(startTime))
	}()

	wg.Wait()

	fmt.Println("Completed in", time.Now().Sub(startTime))
	fmt.Println("Bye")
}
