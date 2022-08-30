package main

import "fmt"

func main() {

	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke-", counter)
		counter++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan Singkat ke-", i)
	}
}
