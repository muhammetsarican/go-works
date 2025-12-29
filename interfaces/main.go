package main

import "fmt"

func main() {
	fmt.Println("interfaces")

	book := book{desi: 3}
	electronic := electronic{desi: 2}

	fmt.Println(book)
	fmt.Println(electronic)

	fmt.Println(book.shippingCost())
	fmt.Println(electronic.shippingCost())

	var product IShippable

	product = book
	fmt.Println(product.shippingCost())

	product = electronic
	fmt.Println(product.shippingCost())

	fmt.Println(calculateShippingCost([]IShippable{book, electronic}))

	var encoder IEncoder
	encoder = &xEncoder{}
	encoder.encode()
	encoder.decode()

	encoder = &yEncoder{}
	encoder.encode()
	encoder.decode()
}

// interface

type IShippable interface {
	shippingCost() int
}

type book struct {
	desi int
}

type electronic struct {
	desi int
}

func (b book) shippingCost() int {
	return b.desi * 10
}

func (e electronic) shippingCost() int {
	return e.desi * 20
}

func calculateShippingCost(products []IShippable) int {
	var total int
	for _, item := range products {
		total += item.shippingCost()
	}
	return total
}

type IEncoder interface {
	encode()
	decode()
}

type xEncoder struct {
}

type yEncoder struct {
}

func (x *xEncoder) encode() {
	fmt.Println("Encoded with xEncoder")
}

func (x *xEncoder) decode() {
	fmt.Println("Decoded with xEncoder")
}

func (y *yEncoder) encode() {
	fmt.Println("Encoded with yEncoder")
}

func (y *yEncoder) decode() {
	fmt.Println("Decoded with yEncoder")
}
