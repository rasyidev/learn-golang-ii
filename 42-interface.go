package main

import "fmt"

/**
- Interface adalah tipe data abstrak, tidak memiliki implementasi
  langsung.
- Interface berisi definisi - definisi method
- Biasanya digunakan sebagai kontrak
- Saat menggunakan interface, wajib menggunakan semua method yang ada
  pada interface tersebut
*/

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return "Halo, saya " + person.Name
}

type Kucing struct {
	Name, Warna string
}

func (kucing Kucing) GetName() string {
	return "Helaw paw lovers, aku " + kucing.Name + ". warna buluku " + kucing.Warna
}

func sayHello(hasName HasName) {
	fmt.Println(hasName.GetName())
}

func main() {
	rasyidev := Person{Name: "Rasyidev"}
	olene := Kucing{
		Name:  "Olene",
		Warna: "Orange",
	}

	// Person dan Kucing harus implement GetName() nya HasName kalo mau pake sayHello()
	sayHello(rasyidev)
	sayHello(olene)
}
