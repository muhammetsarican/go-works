package main

import "fmt"

func main() {
	fmt.Println("functions")
	add(1, 2)

	total := addWithReturn(1, 2)
	fmt.Println(total)

	total, diff := addWithTwoReturns(1, 2)
	fmt.Println(total, diff)

	total, diff, product := calculationWithThreeReturns(1, 2)
	fmt.Println(total, diff, product)

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(totalOfSlice(numbers))

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func add(a int, b int) {
	fmt.Println(a + b)
}

func addWithReturn(a int, b int) int {
	return a + b
}

func addWithTwoReturns(a int, b int) (int, int) {
	return a + b, a - b
}

func calculationWithThreeReturns(a int, b int) (int, int, int) {
	return a + b, a - b, a * b
}

func totalOfSlice(slice []int) int {
	total := 0
	for _, value := range slice {
		total += value
	}
	return total
}

// using args
func sum(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}
	return total
}
