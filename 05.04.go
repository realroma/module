package main

import (
	"fmt"
	"time"
	"sync"
)

func even(arr [5]int, wg *sync.WaitGroup, ch chan<- int, oddCh chan<- int, evenCh <-chan int) {
	fmt.Println("Gorutine even is srtarting.")
	for i := range arr {
		a := <-evenCh
		if a == 0 {
			oddCh <- 0
			ch <- arr[i]
		}
	}
	wg.Done()
}

func odd(arr [5]int, wg *sync.WaitGroup, ch chan<- int, oddCh <-chan int, evenCh chan<- int) {
	fmt.Println("Gorutine odd is starting.")
	for i := range arr {
		b := <-oddCh
		if b == 0 {
			evenCh <- 0
			ch <- arr[i]
		}
	}
	wg.Done()
}

func add(ch chan int, arr *[10]int) {
	fmt.Println("Add in queue.")
	a := 0
	for i := range ch {
		arr[a] = i
		a++
	}
}

func main() {
	var evenArr [5]int = [5]int{1, 3, 5, 7, 9}
	var oddArr [5]int = [5]int{2, 4, 6, 8, 10}
	var arr [10]int
	fmt.Println("Creating chanels.")
	ch := make(chan int, 10)
	evenCh := make(chan int, 1)
	oddCh := make(chan int, 1)
	evenCh <- 0
	wg := &sync.WaitGroup{}
	fmt.Println("Starting gorutines.")
	wg.Add(2)
	go even(evenArr, wg, ch, oddCh, evenCh)
	go odd(oddArr, wg, ch, oddCh, evenCh)
	wg.Wait()
	close(ch)
	go add(ch, &arr)
	time.Sleep(1 * time.Second)
	fmt.Println(arr)
}
