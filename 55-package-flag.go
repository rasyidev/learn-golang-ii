package main

import (
	"flag"
	"fmt"
	"learn-go-ii/helper"
)

/*
Package flag digunakan untuk membuat input untuk aplikasi berbasis terminal
*/

func main() {
	makanan := flag.String("makanan", "nasi", "Masukkan makanan favorit")
	minuman := flag.String("minuman", "air putih", "Masukkan makanan favorit")
	bayar := flag.Float64("bayar", 0, "Masukkan nominal pembayaran")

	flag.Parse()
	fmt.Println("Terima kasih sudah membeli", *makanan, "dan", *minuman, "kami. \nHave a nice day!")
	helper.Separate("Total Pembayaran", *bayar)
}

//Input: go run 55-package-flag.go -makanan "Bakso" -minuman "Kelapa muda" -bayar 40000
