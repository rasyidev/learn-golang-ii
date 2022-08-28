package main

import "fmt"

func main() {
	// var <nama variabel> <tipe data (wajib)>
	var name string
	name = "Rasyidev"

	fmt.Println("Halo nama saya", name)

	name = "Rasyidev Pro"
	fmt.Println("Nama baru saya", name)


	// var <nama variabel> <tipe data (optional)> := <value>
	var isHappyToLearn bool = true
	var isLearningGo = true

	fmt.Println("Is Rasyidev Happy to Learn Go:", isHappyToLearn)
	fmt.Println("Is Rasyidev Learning Go:", isLearningGo)
}