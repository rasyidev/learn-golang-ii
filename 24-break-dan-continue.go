package main

import "fmt"

func main() {
	student := make(map[string]string)
	student["name"] = "Rasyidev Pro"
	student["nim"] = "2323532"
	student["prodi"] = "Sistem Informasi"

	for key, value := range student {
		if key == "nim" {
			fmt.Println("key = nim | Perulangan dihentikan")
			break
		}
		fmt.Println(key, ":", value)
	}

	fmt.Println("--------------------------------------------------------")

	for key, value := range student {
		if key == "nim" {
			fmt.Println("key = nim | Perulangan pada bagian ini di-skip")
			continue
		}
		fmt.Println(key, ":", value)
	}

}
