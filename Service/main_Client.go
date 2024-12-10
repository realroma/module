package main

import (
	"fmt"
	"net/http"
	"strings"
)

// func ServerRead(SWrite http.ResponseWriter, request *http.Request) {
// 	fmt.Fprintf(request)
// }

func main() {
	r := strings.NewReader("my request")
	resp, _ := http.Get("http://localhost:8080", "/metrics", r)
	defer resp.Body.Close()
	fmt.Println("Some")
}
