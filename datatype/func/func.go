package main

import "fmt"

func someF1(b int) (val int) {
	//Val Инициализируется при входе.
	if b == 2 {
		val = 2
	} else if b == 1 {
		val = 1
	}
	return
}

// Принимает копию объекта val.
//	Равносильно &val что передаёт значение ячейки.
//	Но если использовать *val то будет меняться значение переданной ячейки.

type somestruct struct {
	field int
}

func someF2(val int) int {
	if val == 2 {
		val = 4
	} else if val == 1 {
		val = 2
	}
	//
	return val
}

// Мы можем возвращать функции в return. Надо дописать.
func funin() func(a, b int) (c int) { //Здесь мы записываем функцию в любую переменную. То есть мы можем потом предать в переменную значения и они выполнятся.
	return func(a, b int) (c int) { return a + b } //a,b int значит что эти две переменных одного типа.
}

func main() {
	//Анонимная функция - одноразовая функция делающая некоторые действия без присвоения в переменную.
	f := func(x, y int) int { return x + y }
	fmt.Println(f(1, 2))

	fmt.Printf("Somefunc: \n1:%v,\n2:%v,\n3:%v", someF1(1), someF1(2), someF1(3))
	fmt.Printf("Somefunc: \n1:%v,\n2:%v,\n3:%v", someF2(1), someF2(2), someF2(3))

	//Некоторые типы, создаваемые в коде, потом не используются в функциях, выдавая ошибку. Здесь всё нормально, но в другом месте была ошибка именно с этим. Надо понять почему.
	type somestruct struct {
		field  int
		field2 int
	}
	var s somestruct
	testStruct := func(s somestruct) { fmt.Println(s) }
	testStruct(s)

	a := 1
	b := 2
	c := funin()         //Пример функции в переменной.
	fmt.Println(c(a, b)) //Выполняем переменную.
}
