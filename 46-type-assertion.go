package main

import "fmt"

type Any interface{}

func main() {
	var industriEra Any = 4.0
	fmt.Println(industriEra)

	result := industriEra.(float64)
	fmt.Println(result)

}
