package main

import (
	"fmt"
	"reflect"
)

// Объявление переменной как типа:
type f int

type struc struct {
	field1 int
	field2 string
	//Структура с тегом
	field3 string `example:"db"`
}

// Metod for structure.
func (s *struc) function(i int, str string) {
	s.field1 = i
	s.field2 = str
}

func main() {
	s := struc{0, "", ""}
	s.function(1, "some")
	//Выводится значения, но сами поля мы не получим.
	fmt.Println(s)

	//По этому мы можем:
	t := reflect.TypeOf(s)

	// Перебираем поля структуры
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println("Имя поля:", field.Name)
		fmt.Println("Тип поля:", field.Type)
		fmt.Println("Тег поля:", field.Tag) // Теги, например, для database, json
	}
}
