package main

import (
	"fmt"
)

func main() {

	v := `Some`
	v = `some`
	fmt.Println(v)

	//Любая строка это массив байтов - это надо помнить.
	a := []byte(v)
	fmt.Println(a)
}
