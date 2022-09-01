package main

import "fmt"

type Pekerja struct {
	Name, Sex string
}

func buatPanggilan(p *Pekerja) {
	if p.Sex == "female" {
		p.Name = "Mbak " + p.Name
	} else {
		p.Name = "Mas " + p.Name
	}
}

func main() {
	linda := Pekerja{
		Name: "Linda",
		Sex:  "female",
	}
	fmt.Println("Linda \n", linda)

	buatPanggilan(&linda)
	fmt.Println("Panggilan Buat Linda\n", linda)
}
