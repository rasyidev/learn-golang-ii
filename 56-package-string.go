package main

import (
	"fmt"
	"learn-go-ii/helper"
	"strings"
)

/*
Package string digunakan untuk memanipulasi string
*/

func main() {
	message := "Rasyidev is learning Go Language"

	message = strings.ToUpper(message)
	helper.Separate("Pesan dalam kapital", message)

	message = strings.ToLower(message)
	helper.Separate("Pesan dalam huruf kecil", message)

	splittedMessage := strings.Split(message, " ")
	for _, val := range splittedMessage {
		fmt.Println(val)
	}

}
