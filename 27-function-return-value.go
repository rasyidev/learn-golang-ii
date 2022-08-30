package main

import "fmt"

func jumlahkan(x float64, y float64) float64 {
	return x + y
}

func main() {
	res := jumlahkan(4, 5)
	fmt.Println("Hasil Penjumlahan:", res)
}
