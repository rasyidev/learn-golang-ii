package main

import "fmt"

func main() {
	a := true
	b := false

	fmt.Println("a=", a, "b=", b)
	fmt.Println("a && b", a && b)
	fmt.Println("a || b", a || b)
	fmt.Println("!a", !a)
	fmt.Println("!b", !b)
}