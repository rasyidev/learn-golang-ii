package main

import "fmt"

func main() {
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{6, 7, 8, 9, 10}

	fmt.Println(iniArray)
	fmt.Println("Array length\t:", len(iniArray))
	fmt.Println("Array capacity\t:", cap(iniArray))
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println(iniSlice)
	fmt.Println("Array length\t:", len(iniSlice))
	fmt.Println("Array capacity\t:", cap(iniSlice))

}
