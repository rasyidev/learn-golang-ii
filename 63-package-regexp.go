package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex regexp.Regexp = *regexp.MustCompile("\\d{2}:\\d{2}:\\d{2}")
	fmt.Println(regex.MatchString("23:24:09"))
	fmt.Println(regex.MatchString("23:24:ee"))
	fmt.Println("---------------------------------------------")

	waktu := "23:02:56 Participant1: Selamat pagi semua\n 53:23:03"
	fmt.Println(regex.FindAllString(waktu, -1)) // -1: semuanya
}
