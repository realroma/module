package main

import (
	"fmt"
)

func main() {
	var m map[string]int
	fmt.Println(m)
	//Пустая карта без значений.
	fmt.Println(m["word"])
	//Получаем значение по ключу, но оно будет пустым
	//Пустые значения для типов string - ""[[]byte, []rune], [int, uint, rune, byte] - 0, [slice, array] - [], struct - {DefaultOfField...}
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
