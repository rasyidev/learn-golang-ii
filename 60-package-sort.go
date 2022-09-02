package main

import (
	"learn-go-ii/helper"
	"sort"
)

/*
- Helper sort digunakan untuk proses pengurutan
- Harus implement interface package sort: Less, Swap, Len
*/

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := UserSlice{
		{"Amri", 24},
		{"John", 20},
		{"Andien", 30},
		{"Benny", 28},
	}
	helper.Separate("Sebelum diurutkan", users)

	sort.Sort(users)
	helper.Separate("Setelah diurutkan", users)

}
