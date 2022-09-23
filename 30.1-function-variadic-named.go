package main

import "fmt"

func sumNamed(numbers ...int64) (result int64) {
	for _, number := range numbers {
		result += number
	}

	return result
}

func main() {
	res := sumNamed([]int64{2, 3, 4, 2, 1}...)
	fmt.Println(res)
	res = sumNamed(2, 4, 5, 3, 1, 5, 6, 2, 1)
	fmt.Println(res)
}
