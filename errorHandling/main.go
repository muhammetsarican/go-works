package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("error handling")

	var a int
	var b int

	var p *int

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(p)

	if p == nil {
		fmt.Println("p is nil")
	}

	// error handling
	file, err := os.ReadFile("test.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))

	result, err := divide(10, 0)

	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(result)

	productService := &ProductService{}

	result, err = productService.CalculatePrice(&Product{
		name:  "Product 1",
		price: 10,
		stock: 1,
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	fmt.Println(
		ValidationError{
			field: "stock",
			error: "stock is not available",
		})
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

type Product struct {
	name  string
	price int
	stock int
}

type ProductService struct {
}

func (p *ProductService) CalculatePrice(product *Product) (int, error) {
	if product.stock < 1 {
		return 0, errors.New("stock is not available")
	}
	if product.price < 0 {
		return 0, errors.New("price is not available")
	}

	return product.price * product.stock, nil
}

type ValidationError struct {
	field string
	error string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("Error: %s, Field: %s", v.error, v.field)
}
