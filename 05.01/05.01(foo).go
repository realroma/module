package main

import "fmt"

func foo() error{
	panic("empty func")
	return nil
}

func main() {
	defer func () {
		err := recover()//Обработка паники и возрат ошибки.
		fmt.Println("Found error:")
		fmt.Println(err)
	}()//Обязательно две скобки в конце () при наличии recover()
	foo()
	fmt.Println("Suckcess")
}
