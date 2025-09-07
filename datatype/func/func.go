package main

import "fmt"

func someF1(b int) (val int) {
	//Val Инициализируется при входе.
	switch {
	case b == 2:
		val = 2
	case b == 1:
		val = 1
	}
	return val
}

func someF1WithPointerOnWalueAndReturnInArguments(v *int) (returnedValue int) {
	fmt.Println(&v) //Адрес ссылки
	*v = 3
	returnedValue = 5
	return
}

// Принимает копию объекта val.
//	Равносильно &val что передаёт значение ячейки.
//	Но если использовать *val то будет меняться значение переданной ячейки.

type somestruct struct {
	field int
}

func someF2(val int) int { //val int - получаем копию значения и работаем с ней. Для получения оригинала нужно получить указатель на данные(*)
	switch {
	case val == 2:
		val = 4
	case val == 1:
		val = 2
	}
	return val
}

// Мы можем возвращать функции в return.
func funin() func(a, b int) (c int) { //Здесь мы записываем функцию в любую переменную. То есть мы можем потом предать в переменную значения и они выполнятся.
	return func(a, b int) (c int) { return a + b } //a,b int значит что эти две переменных одного типа.
}

func ManyArgyments(a ...int) {} //Мы можем принять разное количество переменных в функцию. (Создаётся слайс[]Type(В данном случае []int) а "..." распаковывает.)

// ...

func AnyArguments(slise []int, a ...int) {}

//...

func main() {
	//Анонимная функция - функция присваиваемая в переменную
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

	var v = 5
	fmt.Println(someF1WithPointerOnWalueAndReturnInArguments(&v))

	a := 1
	b := 2
	c := funin()         //Пример функции в переменной.
	fmt.Println(c(a, b)) //Выполняем переменную.

	//Мы объявляем анонимную функицю вот так func(){}, но что тоогда значат ()(круглые скобки) после {}(Фигурных скобок)?
	//Это - вызов анонимной функции сразу.
	func() {
		fmt.Println("Анонимная функция выполнена.")
	}()
}
