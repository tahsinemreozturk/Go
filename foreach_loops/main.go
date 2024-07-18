package main

import (
	"fmt"
)

func main() {
	/*
			var numbers = []int{1, 2, 3, 4}

			for index := 0; index < len(numbers); index++ {
				fmt.Println(numbers[index])
			}

			//index degeri istenirse
			for index, value := range numbers {
				fmt.Println(index, value)
			}

			//index degeri olmadiginda
			for _, value := range numbers {
				fmt.Println(value)
			}

		var language = "Golang"

		for _, character := range language {
			fmt.Printf("Character %c\n", character)
		}*/

	var names = map[string]int{
		"Tahsin ": 1,
		"Emre":    2,
	}

	for key, value := range names {
		fmt.Printf("Key: %s , Value: %d\n", key, value)
	}

}
