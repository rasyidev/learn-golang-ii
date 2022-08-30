package main

import "fmt"

func main() {
	student := make(map[string]string)
	student["name"] = "Rasyidev Pro"
	student["nim"] = "2323532"
	student["prodi"] = "Sistem Informasi"

	for key, value := range student {
		fmt.Println(key, ":", value)
	}
}
