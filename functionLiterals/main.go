package main

import "fmt"

func main() {
	fmt.Println("function literals")

	func() {
		fmt.Println("function literals inner function")
	}()

	func(x int, y int) {
		fmt.Println(x + y)
	}(1, 2)

	var total = func(x int, y int) int {
		return x + y
	}(3, 5)

	fmt.Println(total)

	multiply := func(x int, y int) int {
		return x * y
	}

	divide := func(x int, y int) int {
		return x / y
	}

	fmt.Println(multipleProcess(10, 5, multiply, divide))
}

func multipleProcess(x int, y int, f1 func(a int, b int) int, f2 func(a int, b int) int) (int, int) {
	return f1(x, y), f2(x, y)
}
