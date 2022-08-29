package main

import "fmt"

func main() {
	// array bulan
	var bulan = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var firstTriwulan = bulan[:3]
	fmt.Println(firstTriwulan)

	/*
		data pada slice dan array saling terhubung
	*/
	firstTriwulan[2] = "March"
	fmt.Println("bulan setelah (Maret diubah jadi March)\n:", bulan)
	firstTriwulan = append(firstTriwulan, "Ramadhan")
	fmt.Println("First triwulan + ramadhan:", firstTriwulan)

	fmt.Println("--------------------------------------------------------------------")

	/*
		data antara slice dan array terputus apabila
		terjadi penambahan (append) yang melebihi kapasitas array
	*/
	var bulanMasehiDanHijriah = append(bulan[:], "Muharram", "Rajab")
	bulan[2] = "Maret"

	fmt.Println("bulan:\n", bulan)
	fmt.Println("Bulan masehi dan hijriah:\n", bulanMasehiDanHijriah)
}
