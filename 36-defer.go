package main

import "fmt"

func cobaDefer() {
	fmt.Println(`
		Ini adalah function yang akan dijalankan terakhiran 
		dalam satu block program, walaupun ada error.
	`)
}

func pembagian(a, b float64) float64 {
	if b == 0 {
		panic("Eits, ga boleh bagi pake nol yaa!")
	}
	return a / b
}

func main() {
	defer cobaDefer()
	fmt.Println(pembagian(10, 0)) // error pembagian dengan nol.

	fmt.Println("Hi guys!, Rasyidev di sini.")
}
