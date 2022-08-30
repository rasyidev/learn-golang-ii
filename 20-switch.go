package main

import "fmt"

func main() {
	bulan := 11
	var result string

	switch bulan {
	case 1:
		result = "Januari"
	case 2:
		result = "Februari"
	case 3:
		result = "Maret"
	case 4:
		result = "April"
	case 5:
		result = "Mei"
	case 6:
		result = "Juni"
	case 7:
		result = "Juli"
	case 8:
		result = "Agustus"
	case 9:
		result = "September"
	default:
		result = "Unknown"
	}

	fmt.Println("Angka \t:", bulan)
	fmt.Println("Result \t:", result)

}
