package main

import (
	"fmt"
	"time"
)

var i int

func task(num int) {
	fmt.Printf("Start N %v\n", num)
	var calc int
	for i := 0; i < 1000; i++ {
		calc = i * num
	}
	fmt.Printf("End N %v: calc = %v\n", num, calc)
}

func main() {
	for i := 1; i <= 5; i++ {
		go task(i)
	}
	time.Sleep(100 * time.Millisecond)
}
