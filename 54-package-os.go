package main

/*
Package os digunakan untuk operasi2 yang dapat dilakukan
pada Operating System seperti membuat file, menampilkan working directory,
menagkap arguments dll.
*/

import (
	"learn-go-ii/helper"
	"os"
)

func main() {
	wd, _ := os.Getwd()
	helper.Separate("Working Directory", wd)

	args := os.Args
	helper.Separate("Arguments", args)
}
