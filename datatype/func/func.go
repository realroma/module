package main

import "fmt"

func main() {
	//Анонимная функция - одноразовая функция делающая некоторые действия без присвоения в переменную.

	f := func(x, y int) int { return x + y }
	fmt.Println(f(1, 2))
}
