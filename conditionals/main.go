package main

import "fmt"

func main() {

	var age = 19

	if age >= 18 {
		fmt.Println("Uygun")
	} else {
		fmt.Println("Uygun degil")
	}

	a := 10
	b := 20
	c := 30

	if a >= b && a >= c {
		fmt.Println("A en buyuk")
	} else if b >= a && b >= c {
		fmt.Println("B en buyuk")
	} else {
		fmt.Println("C en buyuk")
	}

}
