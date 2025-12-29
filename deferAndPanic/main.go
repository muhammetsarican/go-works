package main

import "fmt"

func main() {
	defer fmt.Println("deferred")
	fmt.Println("hello")

	// defer runs as stack logic
	defer fmt.Println("deferred 1")
	defer fmt.Println("deferred 2")
	fmt.Println("hello 2")

	// when variable values changes, defer writes old value
	x := 5
	y := 10

	defer fmt.Println(x, y)

	x = 20
	y = 40

	// when program break or something goes wrong, defer will not run

	defer fmt.Println("deferred 3")
	if true {
		// return
	}
	defer fmt.Println("deferred 4")

	// panic
	// panic("something went wrong")

	// when program panics, the defer defined before panic will run
	defer deferRuns()
	panic("something went wrong")

}

func deferRuns() {
	fmt.Println("deferred 5")
}
