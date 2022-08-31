package main

import "fmt"

/**
Ini adalah komentar multiline, biasanya ada di atas function
dan digunakan untuk dokumentasi function. Sedangkan komentar
single-line biasanya digunakan untuk dokumentasi proses di dalam function.
*/
func main() {

	// anonymous function untuk menyapa coders di seluruh dunia
	func() {
		fmt.Println("Helo world, Rasyidev di sini.")
	}()
}
