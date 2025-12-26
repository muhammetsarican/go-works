package main

import (
	"fmt"
	"reflect"
)

func main() {
	var productName string
	var quantity int
	var discount float32
	var isAvailable bool

	productName = "Golang Lectures"
	quantity = 10
	discount = 10.5
	isAvailable = true

	fmt.Println(productName, reflect.TypeOf(productName))
	fmt.Println(quantity, reflect.TypeOf(quantity))
	fmt.Println(discount, reflect.TypeOf(discount))
	fmt.Println(isAvailable, reflect.TypeOf(isAvailable))

	const lectureName = "Golang Lectures"
	const lectureQuantity = 10
	const lectureDiscount = 10.5
	const lectureIsAvailable = true

	fmt.Println(lectureName, reflect.TypeOf(lectureName))
	fmt.Println(lectureQuantity, reflect.TypeOf(lectureQuantity))
	fmt.Println(lectureDiscount, reflect.TypeOf(lectureDiscount))
	fmt.Println(lectureIsAvailable, reflect.TypeOf(lectureIsAvailable))

	orderName := "Golang Lectures"
	orderQuantity := 10
	orderDiscount := 10.5
	orderIsAvailable := true

	fmt.Println(orderName, reflect.TypeOf(orderName))
	fmt.Println(orderQuantity, reflect.TypeOf(orderQuantity))
	fmt.Println(orderDiscount, reflect.TypeOf(orderDiscount))
	fmt.Println(orderIsAvailable, reflect.TypeOf(orderIsAvailable))

	var (
		goName        = "Golang gos"
		goQuantity    = 10
		goDiscount    = 10.5
		goIsAvailable = true
	)

	fmt.Println(goName, reflect.TypeOf(goName))
	fmt.Println(goQuantity, reflect.TypeOf(goQuantity))
	fmt.Println(goDiscount, reflect.TypeOf(goDiscount))
	fmt.Println(goIsAvailable, reflect.TypeOf(goIsAvailable))

	// ids must be int64
	var id int64 = 1
	fmt.Println(id, reflect.TypeOf(id))

	// for more var types format see https://pkg.go.dev/fmt

	// normal print
	fmt.Println("goName:", goName)
	fmt.Println("goQuantity:", goQuantity)
	fmt.Println("goDiscount:", goDiscount)
	fmt.Println("goIsAvailable:", goIsAvailable)

	// formatted print
	fmt.Printf("goName: %s\n", goName)
	fmt.Printf("goQuantity: %d\n", goQuantity)
	fmt.Printf("goDiscount: %f\n", goDiscount)
	fmt.Printf("goIsAvailable: %t\n", goIsAvailable)

	// %v is default format
	fmt.Printf("goName: %v\n", goName)
	fmt.Printf("goQuantity: %v\n", goQuantity)
	fmt.Printf("goDiscount: %v\n", goDiscount)
	fmt.Printf("goIsAvailable: %v\n", goIsAvailable)
}
