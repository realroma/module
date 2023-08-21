package main

import (
	"bufio"
//	"fmt"
	"module/server"
	"net/http"
	"os"
	//	"strings"
//	"io"
)

func main() {
	Server()
	server.Httpserver()
}

func post(w http.ResponseWriter, req *http.Request) {
	p := make([]byte, 1024)
	bufio.NewReader(req.Body).Read(p)
	file, _ := os.Create("log.txt")
	defer file.Close()
	file.Write(p)
	_ = file.Sync()
}

func Server() {
	http.HandleFunc("/", post)
}
