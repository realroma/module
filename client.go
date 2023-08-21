//package clients
package main

import (
	"fmt"
	"net"
)

func /*TCPClient() */main() {
//	p := make([]byte, 1024)
	conn, _ := net.Dial("tcp", "localhost:1234")
	fmt.Fprintf(conn, "Hi!")
	conn.Close()
}
