package main

import (
	"fmt"
	"time"
)

func write(ch chan<- int) {
	ch <- 1
}

func main() {
	ch := make(chan int)
	go write(ch)
	time.Sleep(1 * time.Second)
	fmt.Println(ch)
	close(ch)
}

