package main

import "fmt"

func main() {
	var buah = [5]string{
		"Apel",
		"Nangka",
		"Manggis",
		"Mangga",
	}

	buah[2] = "Pisang"

	fmt.Println(buah)
	fmt.Println("Total buah \t:", len(buah))
}
