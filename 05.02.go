package main

import (
	"errors"
	"fmt"
	"time"
)

func write(ch chan<- int) {
	for i := 0; i <= 3; i++ { //Записываем в небуферизированный канал значения.
		ch <- i
	}
	close(ch) //
}

func job(ch <-chan int, t *time.Timer) {
	quit := make(chan int)
	time.Sleep(3 * time.Second) //
	select {
	case <-t.C:
		err := errors.New("Time is out.")
		fmt.Println(err)
	case <-quit:
		return
	case <-ch:
		for i := range ch { //
			fmt.Printf("Done %v.\n", i)
		}
		if !t.Stop() { //
			<-t.C
		}
	}
}

func main() {
	timer := time.NewTimer(1 * time.Second) //
	ch := make(chan int)                    //
	go write(ch)                            //

	go job(ch, timer)

	time.Sleep(3 * time.Second) //
	fmt.Println("Done")
}
