package main

import "fmt"

func main() {
	var age = 17

	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can't vote")
	}

	a := 10
	b := 20
	c := 30

	if a > b && a > c {
		fmt.Println("a is the largest")
	} else if b > a && b > c {
		fmt.Println("b is the largest")
	} else {
		fmt.Println("c is the largest")
	}
}
