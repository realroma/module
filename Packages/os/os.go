package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("Test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}
