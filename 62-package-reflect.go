package main

import (
	"learn-go-ii/helper"
	"reflect"
)

type Sample struct {
	Name string
}

func main() {
	sample := Sample{"Rasyidev"}

	typeOfSample := reflect.TypeOf(sample)
	helper.Separate("Tipe dari Sample", typeOfSample)
}
