package main

import (
	"fmt"
	"net/http"
)

func ServerRead(SWrite http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(request)
}


func main() {
	resp, _ := http.Get("http://localhost:8080", "/metrics", )
	fmt.Println("Some")
}
