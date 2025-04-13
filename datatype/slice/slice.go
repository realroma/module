package main

import (
	"fmt"
)

// Динамический массив(Слайс, срез) - структура(!) котора имеет вид:
// type slise struct {
// Adress *Pointer - Ссылка на аддрес памяти с массивом данных.
// Capasiti int - Максимально возможный размер до аллокации(перемещения, переприсвоения)
// Lenght int - Заполненность данными. Может отсутствовать наполнение, но не размер.
// }

// Если он ссылается на какой-то массив, то при урезании([value:value]) он может сохранить часть размера из за связности
func main() {
	var a []int = []int{1, 2, 3, 4}
	fmt.Println(a)

	a = a[:0]
	fmt.Printf("Slise: %v,\n Lenght: %v,\n Capasiti: %v\n", a, len(a), cap(a))

	a = a[:2]
	fmt.Printf("Slise: %v,\n Lenght: %v,\n Capasiti: %v\n", a, len(a), cap(a))

	a = a[2:]
	fmt.Printf("Slise: %v,\n Lenght: %v,\n Capasiti: %v\n", a, len(a), cap(a))
}
