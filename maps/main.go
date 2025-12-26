package main

import "fmt"

func main() {
	fmt.Print("maps")

	namesMap := make(map[string]int, 0)

	namesMap["Muhammet"] = 1
	namesMap["Rubypome"] = 2
	namesMap["Amsterdam"] = 3

	fmt.Println(namesMap)
	fmt.Println(namesMap["Muhammet"])

	// add values at creation
	agesMap := map[string]int{
		"Muhammet":  1,
		"Rubypome":  2,
		"Amsterdam": 3,
	}

	fmt.Println(agesMap)
	fmt.Println(agesMap["Muhammet"])

	// delete
	delete(namesMap, "Muhammet")
	fmt.Println(namesMap)
}
