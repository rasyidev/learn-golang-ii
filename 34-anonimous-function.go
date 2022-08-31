package main

import "fmt"

type RegisterStatus func(string) bool

func checkIsRegistered(name string, rs RegisterStatus) {
	res := rs(name)
	fmt.Println("User Register Status\t:", res)
}

func main() {
	isRegistered := func(name string) bool {
		return name == "admin"
	}

	fmt.Println("Admin Register Status\t:", isRegistered("editor"))

	fmt.Println("-------------------------------------------------------------")

	checkIsRegistered("admin", isRegistered)
}
