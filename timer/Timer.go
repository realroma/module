package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now()
	fmt.Println(a)

	//Остановка всего приложения или только main функции?
	time.Sleep(2 * time.Second)

	s := time.Now()

	//Измеряется отсносительно другого времени
	t := s.Sub(a)

	//Создаём отрезок времени, и если он больше определённого времению
	second, _ := time.ParseDuration("5s")
	if t > second {
		fmt.Println("Time: ", t)
	} else {
		fmt.Println("t-cicle less than second")
	}

}
