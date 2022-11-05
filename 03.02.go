package main

import "fmt"

func add(arr [5]int, slise []int) bool {
		for i := range arr {
			for f := range slise {
				y := slise[f]
				if arr[i] == y {
					return true
				}
			}
			slise = append(slise, arr[i])
		}
		return false
}

func main () {
	var arr = [...]int{1,2,3,4,5}
	var arr2 = [...]int{1,1,3,4,4}
	slise := make([]int, 0, 5)
	fmt.Println(add(arr2, slise))
	fmt.Println(add(arr, slise)) 
}
