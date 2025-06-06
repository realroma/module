package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Now()
	fmt.Println(a)
	time.Sleep(2 * time.Second)
	s := time.Now()
	//Измеряется отсносительно другого времени
	time := s.Sub(a)
	fmt.Println(time)
	if time > 1 {
		fmt.Println("Im string!")
	}

}
