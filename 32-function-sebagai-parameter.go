package main

import "fmt"

func filterBrand(text string) string {
	if text == "Indomie" {
		return "Mie Instan"
	}

	return text
}

func sayMakananFavorit(makanan string, filter func(string) string) {
	res := filter(makanan)
	fmt.Printf("Saya suka makan %v", res)
}

func main() {
	sayMakananFavorit("Indomie", filterBrand)
}
