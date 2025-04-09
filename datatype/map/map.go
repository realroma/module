package main

import (
	"fmt"
)

func main() {
	//
	var m map[string]int
	fmt.Println(m)
	//Value for default
	fmt.Println(m["word"])
	//Try check is value in the map
	val, ok := m["word"]

	fmt.Println("Status: ", ok)
	fmt.Println("Value: ", val)

	var map2 map[struct{}][]bool

	m = map[string]int{
		"Some": 1,
	}

	delete(m, "Some")

	fmt.Println(map2)

	//Make() can create only chanel, map, slise.
	m3 := make(map[int]int)
	fmt.Println(m3)
}
