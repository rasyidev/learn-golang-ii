package main

import (
	"fmt"
	"strings"
)

func getPartNames(fullName string) (firstName, middleName, lastName string) {
	name := strings.Split(fullName, " ")
	firstName = name[0]
	middleName = name[1]
	lastName = name[2]
	return
}

func main() {
	a, b, c := getPartNames("Rasyidev Pro Player")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
