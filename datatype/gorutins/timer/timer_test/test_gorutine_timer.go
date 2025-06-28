package main

import (
	"time"
	"fmt"
)

func something(t *time.Timer) {
	time.Sleep(2 * time.Second)
	select {
	case <-t.C:
		fmt.Println("Time is out.")
	default:
		fmt.Println("Fall in default.")
	}
}

func main() {
	timer := time.NewTimer(1 * time.Second)
	go something(timer)
	time.Sleep(3 * time.Second)
	fmt.Println("Main ends.")
}
