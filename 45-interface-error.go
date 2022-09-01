package main

// Go lang sudah include interface error. Bisa kita return tipe data apapun
import (
	"errors"
	"fmt"
)

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Tidak bisa pembagian dengan nol")
	}

	return a / b, nil
}

func main() {
	a, err := div(2, 4)
	fmt.Println(a, err)

	b, err := div(13, 0)
	fmt.Println(b, err)
}
