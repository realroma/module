package main

import (
	"fmt"
	"time"
)

func went(t *time.Timer){
	time.Sleep(2 * time.Second)
	select {
	case <-t.C:
		fmt.Println("Time is out.")
	default:
		fmt.Println("Ready.")
	}
}

func main() {
	timer := time.NewTimer(1 * time.Second)
	go went(timer)
	timer = time.NewTimer(3 * time.Second)
	go went(timer)
	fmt.Println("Wait")
	time.Sleep(4 * time.Second)
	fmt.Println("Main is ends.")
}
