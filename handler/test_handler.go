package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Http header."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
}
