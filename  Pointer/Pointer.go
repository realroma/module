package main

import (
	"fmt"
)

type stru struct {
	f1, f2 string
	f3     int
}

func main() {
	variable := 1
	fmt.Println("Adress in memory:", &variable)
	pointer := &variable
	fmt.Printf("Variable pointer is: %T\n", pointer)
	fmt.Println("Value in pointer", *pointer)

	s := &stru{
		f1: "Sometext",
		f3: 1,
	}
	s.f2 = "Othertext"
	fmt.Printf("%+v\n", s)
	//Один важный факт: для структур можно использовать (*struct).fild и разница с struct.fild будет только в чтении.
	//В Go разименовывание полей для структур автоматическое.
	//То же действенно и в отношении массивов и срезов. Желательно изучить подробнее и провести тесты.
}
