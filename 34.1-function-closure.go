package main

import "fmt"

/*
- function closure: function yang bisa disimpan dalam variabel
- Bisa dipakai untuk membuat function di dalam function
- function closure bentuknya adalah anonymous function
*/

func main() {
	var getMinMax = func(nums []int) (int, int) {
		min := nums[0]
		max := nums[0]

		for _, num := range nums {
			switch {
			case num > max:
				max = num

			case num < min:
				min = num
			}
		}

		return min, max
	}

	numbers := []int{1, 3, 4, 5, 10, 2, 6, 7, 8, 0}
	min, max := getMinMax(numbers)

	fmt.Printf("data\t: %v \nmin\t: %v\nmax\t: %v", numbers, min, max)
}
