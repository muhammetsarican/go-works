package main

import "fmt"

func main() {
	fmt.Println("forEachLoops")

	var numbers = []int{1, 2, 3, 4, 5}

	// foreach arrays
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	for _, value := range numbers {
		fmt.Println(value)
	}

	var codingLanguages = []string{"Go", "Python", "JavaScript", "Java", "C#"}

	for index, value := range codingLanguages {
		fmt.Println(index, value)
	}

	for _, value := range codingLanguages {
		fmt.Println(value)
	}

	// foreach maps
	var languages = map[string]string{
		"en": "English",
		"tr": "Turkish",
		"de": "German",
	}

	for key, value := range languages {
		fmt.Println(key, value)
	}

	for key := range languages {
		fmt.Println(key)
	}

	for value := range languages {
		fmt.Println(value)
	}

	// foreach strings
	var name = "Muhammet"

	for index, value := range name {
		fmt.Println(index, value)
	}

	for _, character := range name {
		fmt.Println(character)
	}
}
