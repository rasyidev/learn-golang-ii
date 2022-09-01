package main

import "fmt"

type Any interface{}

func main() {
	var variabelTipeApapun interface{}

	variabelTipeApapun = 24
	fmt.Println(variabelTipeApapun)

	variabelTipeApapun = []int{2, 4, 5, 2, 5, 6}
	fmt.Println(variabelTipeApapun)

	variabelTipeApapun = "Rasyidev"
	fmt.Println(variabelTipeApapun)

	person := map[string]Any{
		"name": "Rasyidev",
		"age":  22,
	}
	variabelTipeApapun = []Any{1, "Satu", "232.4", person}
	fmt.Println(variabelTipeApapun)
}
