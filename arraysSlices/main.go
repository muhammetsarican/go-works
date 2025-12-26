package main

import "fmt"

func main() {
	var whoAmI [5]string

	whoAmI[0] = "Muhammet"
	whoAmI[1] = "Rubypome"
	whoAmI[2] = "Amsterdam"
	whoAmI[3] = "32"
	whoAmI[4] = "182"

	fmt.Println(whoAmI)

	var names = [5]string{"Muhammet", "Rubypome", "Amsterdam", "32", "182"}
	fmt.Println(names)
	names[1] = "Mehmet"
	fmt.Println(names[0:3])

	var slicedNames = []string{"Muhammet"}
	slicedNames = append(slicedNames, "Mehmet")
	fmt.Println(slicedNames)
}
