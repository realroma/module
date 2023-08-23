package main

import (
	"fmt"
	"flag"
	"net/http"
)

func main() {
	url := flag.String("url", "", "Path to server")
	n := 1
	flag.Parse()
	fmt.Printf("Overall status is %v, with service_id %v mesql component is %v\n", n, n, *url)
}

func HTTPClient() {
	resp, _ := htt.Get
}
