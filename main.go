package main

import "fmt"

const (
	PassStatuse = "pass"
	FailStatuse = "fail"
)

type HealthCheck struct {
	ServisID int
	status string
}

func GenerateCheck() []HealthCheck {
	var slise []HealthCheck //Cоздаём слайс структур типа HealthCheck
	var s HealthCheck
	for i := 0; i<5; i++ { // 0-4
		if 0 == i % 2 { // Каждое чётное число
			s = HealthCheck{i, PassStatuse} //Создаём переменную со структурой
			slise = append(slise, s) //Добавляем в слайс структуры
		} else {
			s = HealthCheck{i, FailStatuse}
			slise = append(slise, s)
		}
        }
	return slise //Возвращаем слайс структур
}

func Run() {

}

func main () {
	a := GenerateCheck() //Присваиваем слайс с структурами в переменную
	for i := range a { //Идёт по индексам структур (№)
		f := a[i] //Присвоение структуры переменной с помощью индекса
		if f.status == PassStatuse { //Проверка статуса допуска
			fmt.Println(f.ServisID)
		}
	}
}
