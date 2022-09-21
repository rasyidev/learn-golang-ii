package tambahan

import (
	"fmt"
	"testing"
)

func TestIfElseVariabelTambahan(t *testing.T) {
	currentYear := 2021

	if age := currentYear - 1999; age < 17 {
		fmt.Println("Belum boleh buat kartu sim")
	} else {
		fmt.Println("Sudah boleh buat kartu sim")
	}
}

/*
=== RUN   TestIfElseVariabelTambahan
Sudah boleh buat kartu sim
--- PASS: TestIfElseVariabelTambahan (0.00s)
PASS
ok      learn-go-ii/tambahan    0.463s
*/
