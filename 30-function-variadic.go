package main

import "fmt"

func sumAll(numbers ...float64) float64 {
	result := 0.0
	for _, number := range numbers {
		result += number
	}
	return result
}

func sumAllBarang(itemName string, prices ...float64) string {
	result := 0.0
	for _, price := range prices {
		result += price
	}

	return fmt.Sprintf("%s: %v", itemName, result)
}

func main() {
	res1 := sumAll(1, 2, 3, 4, 5)
	res2 := sumAllBarang("Kayu", 2, 4, 5, 6, 2.2, 2, 5, 6)

	numbers := []float64{1, 3, 5, 7.3, 7.24, 1.6}
	res3 := sumAll(numbers...)

	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}
