package main

import "fmt"

func lingkaran(r float64) (float64, float64) {
	keliling := 2 * 22 / 7 * r
	luas := 22 / 7 * r * r
	return keliling, luas
}

func main() {
	k, l := lingkaran(24)
	fmt.Println("Keliling Lingkaran \t:", k)
	fmt.Println("Luas Lingkaran \t\t:", l)
}
