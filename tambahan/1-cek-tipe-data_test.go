package tambahan

import (
	"fmt"
	"testing"
)

func TestCekTipeData(t *testing.T) {
	angka := 1
	panjang := 23.4
	nama := "Rasyidev Pro"
	orang := map[string]interface{}{
		"Name": "Rasyidev Pro",
		"Age":  22,
	}

	fmt.Printf("Tipe dari %v \t: %T\n", angka, angka)
	fmt.Printf("Tipe dari %v \t: %T\n", panjang, panjang)
	fmt.Printf("Tipe dari %v \t: %T\n", nama, nama)
	fmt.Printf("Tipe dari %v \t: %T\n", orang, orang)
}

/*
=== RUN   TestCekTipeData
Tipe dari 1     : int
Tipe dari 23.4  : float64
Tipe dari Rasyidev Pro  : string
Tipe dari map[Age:22 Name:Rasyidev Pro]         : map[string]interface {}
--- PASS: TestCekTipeData (0.00s)
PASS
ok      learn-go-ii/tambahan    (cached)
*/
