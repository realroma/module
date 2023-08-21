package main

import (
	"fmt"
	"net/http"
//	"bufio"
	"strings"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n") //просто печатаем нашу строку в ResponseWriter, он поддерживает Writer
}

func HTTPClient() {
	r := strings.NewReader("my request")
	resp, _ := http.Post("http://localhost:8080", "/POST", r)
	defer resp.Body.Close()
}

func main() {
	HTTPClient()
}

