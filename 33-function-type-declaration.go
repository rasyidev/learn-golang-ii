package main

import (
	"fmt"
	"strings"
)

type Filter func(string) string

func upperCaseFilter(text string) string {
	return strings.ToUpper(text)
}

func sayHelloWithFilter(nama string, filter Filter) {
	res := filter(nama)
	fmt.Println("Hi", res, "nice to meet you")
}

func main() {
	sayHelloWithFilter("rasyidev", upperCaseFilter)
}
