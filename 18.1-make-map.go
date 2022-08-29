package main

import "fmt"

func main() {
	student := make(map[string]string)
	student["name"] = "Rasyidev Pro"
	student["nim"] = "2323532"
	student["prodi"] = "Sistem Informasi"

	fmt.Println(student)
	fmt.Println(len(student))
	delete(student, "prodi")

	fmt.Println("prodi removed from student")
	fmt.Println(student)
	fmt.Println(len(student))
}
