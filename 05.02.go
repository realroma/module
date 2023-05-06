package main

import (
	"errors"
	"fmt"
	"time"
)

func write(ch chan<- int) {
	fmt.Println("Writing in chanel")
	for i := 0; i <= 3; i++ { //Записываем в небуферизированный канал значения.
		ch <- i
	}
	close(ch) //
}

func read(ch <-chan int, t *time.Timer) {
	time.Sleep(2 * time.Second)
	fmt.Println("Seting for read")
	select {
	case <-t.C:
		err := errors.New("Time is out.")
		fmt.Println(err)
		return
	case <-ch:
		for i := range ch {
			fmt.Printf("Done %v.\n", i)
		}
	}
}

func MyFunc () {
	ch := make(chan int)
	go write(ch)
	timer := time.NewTimer(1 * time.Second)
	go read(ch, timer)
	time.Sleep(3 * time.Second)
}

func main() {
	fmt.Println("MyFunc is starting")
	MyFunc()
}
