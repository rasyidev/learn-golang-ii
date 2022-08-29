package main

import "fmt"

func main() {
	const PI = 3.14
	r := 21
	luas_lingkaran := PI * float64(r) * float64(r)
	keliling_lingkaran := 2 * PI * float64(r)
	fmt.Println("Luas Lingkaran \t:", luas_lingkaran)
	fmt.Println("Keliling Lingkaran \t:", keliling_lingkaran)

	
	gaji_pertahun := 100800000
	gaji_perbulan := gaji_pertahun / 12
	fmt.Println("Gaji Pertahun\t:", gaji_pertahun)
	fmt.Println("Gaji Perbulan\t:", gaji_perbulan)
}