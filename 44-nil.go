package main

import "fmt"

func main() {
	// default value untuk variabel bukan nil
	var s string                      // ""
	var f float64                     // 0
	var i int64                       // 0
	var slice = []int{}               // []
	m := make(map[string]interface{}) // map[]

	fmt.Println(s, f, i, slice, m)
}
