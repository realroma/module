package main

import (
//	"os"
	"net/http"
	"fmt"
	"io/ioutil"
	"project/module/server"
)

func json(w http.ResponseWriter, req *http.Request) {
	file, _ := ioutil.ReadFile("DataBase/db.json")
	a := string(file)
	fmt.Println(a)
	fmt.Fprintf(w, a)
}

func main() {
	http.HandleFunc("/health", json)
	server.HTTPServer()
}
