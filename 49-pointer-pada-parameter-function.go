package main

import "fmt"

type Orang struct {
	Name, Race string
	Age        int
}

func Javanize(o Orang) {
	o.Race = "Javanese"
}

// Fungsi Bugize cuma bisa nerima pointer objek Orang (pakai &)
func Bugize(o *Orang) {
	o.Race = "Bugis"
}

func main() {
	andien := Orang{
		Name: "Andien",
		Age:  24,
		Race: "Sundanese",
	}

	kepin := Orang{
		Name: "Kepin",
		Age:  25,
		Race: "Malay",
	}

	Javanize(andien)
	fmt.Println(andien)

	Bugize(&kepin)
	fmt.Println(kepin)

}
