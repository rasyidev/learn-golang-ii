package main

import "fmt"

type Any interface{}
type Orang struct {
	Name string
	Age  int
	Race string
}

func main() {
	var a Any = 23

	// b adalah alamat dari a
	b := &a

	// nilai yg ada pada alamat b diubah jadi "Rasyidev"
	*b = "Rasyidev"
	fmt.Println(a, *b)

	andi := Orang{
		Name: "Andi",
		Age:  23,
		Race: "Javanese",
	}

	// andi2 adalah alamat dari andi
	andi2 := &andi
	// andi2 (alamat) diisi dengan (alamat) Orang,
	// jadinya andi2 nilainya diganti Orang baru
	andi2 = &Orang{
		Name: "Kim Tae Ri",
		Age:  31,
		Race: "Korean",
	}

	fmt.Println(andi)
	fmt.Println(andi2)
	fmt.Println("-----------------------------------------------------")

	// andi3 adalah alamat dari andi
	andi3 := &andi
	// nilai dari semua variabel yg menunjuk andi3 (andi juga punya alamat yg sama)
	// akan diubah jadi Orang baru, nilai dari andi dan andi3 sama aja karena menunjuk
	// alamat yg sama
	*andi3 = Orang{
		Name: "Punch",
		Age:  28,
		Race: "Asian",
	}
	fmt.Println(andi)
	fmt.Println(andi2)
	fmt.Println(andi3)

}
