package main

import (
	"fmt"
)

// Строки в golang не изменяемый тип данных. Невозможно никаким способом изменить их кроме конкотенации(сложения).
func main() {

	v := `Some`
	v = `some`
	fmt.Println(v)

	//Любая строка это массив байтов(ASCLL), рун при использовании Unix - это надо помнить.
	a := []byte(v)
	fmt.Println(a)
	fmt.Println([]byte("{\"some\":\"map\"}"))
	//Написать стороку в необработанном свиде:
	fmt.Println(`Strings text, some spase \n, simvols //,|
	//  not work with income value - % g`)

	//Понять почему символы в "!" строки а в '!' руны. Что же тогда будет `!`? Как будут байты?
}
