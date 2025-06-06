package main

import "fmt"

func main() {
	s := "Some String"
	print(s, "\n")

	fmt.Print(s, "\n")

	fmt.Println(s)

	fmt.Printf("%v", s)
}
