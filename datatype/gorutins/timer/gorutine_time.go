package main

import (
	"time"
	"fmt"
)

func main() {
	t0 := time.Now()
	i := 0
	for i < 1000000	 {
		i = i + 1
	}
	fmt.Println(i, time.Since(t0))
}
