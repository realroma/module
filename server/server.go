package server

import (
	"net/http"
	"fmt"
)

func Httpserver() {
	fmt.Println("start")
	http.ListenAndServe(":8080", nil)
}
