package main

import "fmt"

func main() {
	index := 1

	for index <= 10 {
		fmt.Println(index)
		index++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	total := 0
	for i := 1; i <= 10; i++ {
		total += i
	}
	fmt.Println(total)

	idx := 0
	names := [3]string{"Muhammet", "Rubypome", "Amsterdam"}
	for idx < len(names) {
		fmt.Println(names[idx])
		idx++
	}

	// break
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	// continue
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("contiune")
			continue
		}
		if i == 3 || i == 7 {
			fmt.Println("or continue")
			continue
		}
		fmt.Println(i)
	}
}
