package main

import "fmt"

func main() {
	slice := []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
