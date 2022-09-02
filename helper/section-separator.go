package helper

import "fmt"

func Separate(title string, any interface{}) {
	fmt.Println("--------------", title, "--------------")
	fmt.Println(any)
	fmt.Println("\n---------- End of", title, "-----------")
}
