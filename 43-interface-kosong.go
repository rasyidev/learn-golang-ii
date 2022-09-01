package main

import "fmt"

/**
- Interface kosong adalah interface yang gak punya
  method sama sekali. Semua tipe data bisa implement
- Biasanya digunakan untuk membuat function yang
  parameternya bisa menerima semua tipe data
*/

type Apapun interface{}

func sayApapun(a Apapun) Apapun {
	fmt.Println(a)
	return a
}

func main() {
	sayApapun("Hello")
	sayApapun(1)
	sayApapun(3.5)
}
