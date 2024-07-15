package main

import "fmt"

func main() {
	//Arrays
	var names [3]string

	names[0] = "Tahsin"
	names[1] = "Emre"
	names[2] = "Öztürk"

	var names2 = [3]string{"Tahsin", "Emre", "Öztürk"}

	names2[2] = "Öztürk 2"

	fmt.Println(names2)

	fmt.Println(names2[0:2])

	//Slices
	var names3 = []string{"Tahsin", "Emre", "Öztürk"}

	names3 = append(names3, "Sonradan ekleme yapilabiliyor.")

	fmt.Println(names3)

}
