package main

import "fmt"

func main() {
	fmt.Println("pointers")

	var a int = 10
	var p *int = &a

	fmt.Println("address of a", p)  // address of a
	fmt.Println("value of p", *p)   // value of p
	fmt.Println("address of p", &p) // address of p

	// update the value that kept in pointer
	*p = 20
	fmt.Println("value of a", a)
	fmt.Println("value of p", *p)

	first := 10
	second := first
	pointer := &first
	*pointer = 20
	fmt.Println("value of first", first)
	fmt.Println("value of second", second)

	// add funcs
	add12(a)
	fmt.Println("value of a", a)
	add12Pointer(&a)
	fmt.Println("value of a", a)

	// changeValue func
	changeValue([]int{1, 2, 3})
	fmt.Println("value of a", a)
}

// explain the difference between add12 and add12Pointer functions in a comment block
/*
	add12 function creates a copy of the value that sent to it and adds 12 to it
	add12Pointer function adds 12 to the value that sent to it using pointer
*/
// create a function that adds 12 to the value that sent to it
func add12(x int) { // pass by value
	x = x + 12
}

// create a function that adds 12 to the value that sent to it using pointer
func add12Pointer(x *int) { // pass by reference
	*x = *x + 12
}

// write a comment block that explains * and &
/*
	* is used to dereference a pointer
	& is used to get the address of a variable
*/

//  not primitive types not using with pointer

func changeValue(slice []int) {
	slice[0] = 100
}
