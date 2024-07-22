package main

import (
	"fmt"
)

func main() {
	/*add2(10, 5)
	fmt.Println(add2(11, 16))
	fmt.Println(calculation(11, 16))*/

	var numbers = []int{2, 3, 5, 10, 20}
	fmt.Println(sum(numbers))
	fmt.Println(sum2(1, 2, 3))

}

// Gonderilen degisken kadar alir.
func sum2(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum

}

// slice elemanlarini gezerek deger dondurur
func sum(numbers []int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}
	return sum
}

func add(x int, y int) {
	fmt.Println(x + y)
}

func add2(x int, y int) int {
	return x + y
}

// Birden fazla donderme yapabilir.
func calculation(x int, y int) (int, int, int) {
	return x + y, x - y, x * y
}
