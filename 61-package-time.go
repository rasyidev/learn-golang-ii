package main

import (
	"fmt"
	"learn-go-ii/helper"
	"time"
)

func main() {
	timeNow := time.Now()
	helper.Separate("Waktu saat ini", timeNow)

	fmt.Println(time.Date(2022, 8, 22, 2, 24, 42, 4240, time.UTC))

	// time.RFC3339
	layout := "2006-01-02" // harus banget pake template ini, kalo gak ya jadi gak sesuai :(
	parsedTime, _ := time.Parse(layout, "2022-09-14")
	helper.Separate("Parsed time", parsedTime)
}
