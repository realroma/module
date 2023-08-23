package server

import (
	"net/http"
	"fmt"
)

func HTTPServer() {
	fmt.Println("start")
	http.ListenAndServe(":8080", nil)
}
