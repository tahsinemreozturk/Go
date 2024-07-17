package main

import "fmt"

func main() {

	var names map[string]int

	names = make(map[string]int, 0)

	names["Tahsin"] = 1
	names["Emre"] = 2

	fmt.Println(names)
	fmt.Println(names["Emre"])
	fmt.Println(names["Olmayan bir deger ciktisi 0 olacaktir"])

	//2. Kullanimi
	names2 := map[string]int{
		"Tahsin": 1,
		"Emre":   2,
	}

	delete(names2, "Emre")
	fmt.Println(names2)

}
