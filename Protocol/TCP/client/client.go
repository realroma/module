package client

import (
	"fmt"
	"net"
)

func Start() {
	conn, _ := net.Dial("tcp", "localhost:1234")
	fmt.Fprintf(conn, "Hi!")
	conn.Close()
}
