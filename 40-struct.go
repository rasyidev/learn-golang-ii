package main

import "fmt"

/**
- Template data untuk menggabungkan nol atau lebih tipe data (dapat bervariasi)
- Kumpulan dari field
- Prototype data
*/

// Nama struct dan fieldnya menggunakan PascalCase
type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	// deklarasi struct tipe 1
	c := Customer{
		Name:    "Rasyidev",
		Address: "Jl. Raden Saleh, Indonesia",
		Age:     22,
	}

	// deklarasi struct tipe 2
	d := Customer{"Rasyidev1", "Jl. Mataram", 23}
	_ = d

	var e Customer
	e.Name = "Rasyidev2"
	e.Address = "Jl. Patimura"
	e.Age = 24

	fmt.Println(c)
	fmt.Println("Nama\t:", c.Name)
	fmt.Println("Alamat\t:", c.Address)
	fmt.Println("Usia\t:", c.Age)
}
