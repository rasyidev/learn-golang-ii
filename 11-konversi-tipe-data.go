package main

import "fmt"

func main() {
	var beratKendaraan int8 = 127
	beratKendaraan2 := int64(int32(beratKendaraan) + 1000)
	fmt.Println("Berat Kendaraan \t:", beratKendaraan)
	fmt.Println("Berat Kendaraan baru \t:", beratKendaraan2)

	nama := "Habib"
	fmt.Println(nama[0], string(nama[0]))

}