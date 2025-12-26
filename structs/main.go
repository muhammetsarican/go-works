package main

import "fmt"

func main() {
	fmt.Println("structs")

	customer := customer{
		name: "Muhammet Sarican",
		age:  25,
	}

	fmt.Println(customer)
	customer.sayHello()
	customer.sayHelloWithPointer()

	customer.changeName("Osman Dybala")
	fmt.Println(customer)
}

type customer struct {
	name string
	age  int
}

// struct functions
func (c customer) sayHello() {
	fmt.Println("Hello", c.name)
}
func (c *customer) sayHelloWithPointer() {
	fmt.Println("Hello", c.name)
}

// change name function with pointer
func (c *customer) changeName(name string) {
	c.name = name
}
