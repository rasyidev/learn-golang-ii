package main

import "fmt"

func main() {
	// membuat slice dengan 2 data, dan array dengan kapasitas 5 data
	// array disembunyikan
	buah := make([]string, 2, 5)
	buah[0] = "Mangga"
	buah[1] = "Pisang"

	fmt.Println("Buah\t:", buah)
	fmt.Println("Total Buah\t:", len(buah))
	fmt.Println("Kapasitas Buah\t:", cap(buah))
}
