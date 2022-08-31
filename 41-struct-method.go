package main

import "fmt"

type Programmer struct {
	Username  string
	Languages []string
}

func (p Programmer) sayHello() {
	fmt.Println("Hello, I'm", p.Username)
	fmt.Println("I code using", p.Languages)
}

func main() {
	rasyidev := Programmer{
		Username:  "Rasyidev",
		Languages: []string{"Go", "Python", "JavaScript"},
	}

	rasyidev.sayHello()
}
