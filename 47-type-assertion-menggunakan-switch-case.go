package main

import "fmt"

func random() interface{} {
	return "Hello"
}

func main() {
	r := random()
	switch r.(type) {
	case string:
		fmt.Println("String |", r)
	case int:
		fmt.Println("Int |", r)
	default:
		fmt.Println("Unknown |", r)
	}
}
