package server

import (
	"fmt"
	"net/http"
)

func HTTPServer() {
	fmt.Println("start server")
	http.ListenAndServe(":8080", nil)
}
