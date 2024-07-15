package main

import (
	"fmt"
)

func main() {
	/*
		var productName string
		var qunatitiy int
		var discount float32
		var isInStock bool

		productName = "Go variables 1"
		qunatitiy = 20
		discount = 13.45
		isInStock = true

		fmt.Println(productName, reflect.TypeOf(productName))
		fmt.Println(qunatitiy, reflect.TypeOf(qunatitiy))
		fmt.Println(discount, reflect.TypeOf(discount))
		fmt.Println(isInStock, reflect.TypeOf(isInStock))

		/*var productName string = "Go variables 1"
		fmt.Println(productName, reflect.TypeOf(productName))

		productName := "Go variables 1"
		fmt.Println(productName)*/

	var productName string
	var qunatitiy int
	var discount float32
	var isInStock bool

	productName = "Go variables 1"
	qunatitiy = 20
	discount = 13.45
	isInStock = true

	fmt.Println("Product name:", productName, "Quantity:", qunatitiy, "Discount:", discount, "Is in Stock ?", isInStock)
	fmt.Printf("Prodcut name %s, Quantity %d , Discount: %f , Is in Stock %t\n ", productName, qunatitiy, discount, isInStock)
	fmt.Printf("Prodcut name %v, Quantity %v , Discount: %v , Is in Stock %v\n ", productName, qunatitiy, discount, isInStock)
}
