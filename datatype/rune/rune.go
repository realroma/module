package main

import (
	"fmt"
)

func main() {
	//Руны это значения символов. Почему не байты? Так как байты имеют размер в 8 бит, то символы других языков не могут быть записаны из-за недостатка места.
	//Любая строка - массив битов, это надо помнить.
	var i string = "1"
	v := []rune(i)
	fmt.Println(v)

	var t string = "И"
	v = []rune(t)
	fmt.Println(v)

	var d rune
	fmt.Println(d)
}
