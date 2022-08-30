package main

import "fmt"

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}

func main() {
	ucapanPerpisahan := sayGoodbye

	ucapanPerpisahan("Kim Tae Ri")

}
