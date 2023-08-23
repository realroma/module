package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:1234")
	fmt.Fprintf(conn, "Hi!")
	conn.Close()
}
