package main

import (
	"learn-go-ii/helper"
	"math"
)

func main() {
	height := 23.4208
	height = math.Round(height)
	helper.Separate("Nilai hasil pembulatan", height)

	nilaiUjianMin := math.Min(24, 50)
	helper.Separate("24 vs 50", nilaiUjianMin)
}
