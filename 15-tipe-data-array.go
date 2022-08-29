package main

import "fmt"

func main() {
	var buah [5]string
	buah[0] = "Apel"
	buah[1] = "Durian"
	buah[2] = "Nangka"
	buah[3] = "Pisang"
	buah[4] = "Manggis"

	fmt.Println(buah)
	buahFavorit := buah[3]
	fmt.Println("Buah favorit saya adalah", buahFavorit)
	fmt.Println("Total Buah \t:", len(buah))
	fmt.Println("Kapasitas Buah \t:", cap(buah))

}
