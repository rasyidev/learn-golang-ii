package main

import "fmt"

func main() {
	student := map[string]string{
		"nama": "Rasyidev Pro",
		"nim":  "249909",
	}
	student["prodi"] = "Teknik Biomedis"

	fmt.Println(student)
	fmt.Println(student["nama"])
	fmt.Println(student["nim"])
	fmt.Println(student["prodi"])

}
