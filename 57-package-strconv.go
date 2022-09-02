package main

import (
	"learn-go-ii/helper"
	"strconv"
)

func main() {

	// ubah string to Any: Parse<Any>
	nim := "992830"
	nimInt, _ := strconv.ParseFloat(nim, 64)
	helper.Separate("NIM dalam Integer", nimInt)

	// ubah Any to string: Format<Any>
	noHP := 824802320384
	noHPString := strconv.FormatInt(int64(noHP), 10)
	noHPString = "+62" + noHPString
	helper.Separate("No HP dalam String", noHPString)
}
