package main

import (
	"container/list"
	"fmt"
)

/*
Double linked list
*/

func main() {
	list := list.New()
	list.PushBack("Rasyidev")
	list.PushBack("Kim Tae Ri")
	list.PushFront("Im Yoona")

	fmt.Println(list)

	for elm := list.Front(); elm != nil; elm = elm.Next() {
		fmt.Println(elm.Value)
	}
}
