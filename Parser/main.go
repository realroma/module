package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Parse(path string) io.Reader {
	if path == "" {
		path = "module/Parser/test.txt"
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func main() {
	str := "god=dgsdf"
	fmt.Printf("%q\n", strings.Split(str, "="))

	file := Parse("")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		fmt.Printf("%q\n", strings.Split(str, "="))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
