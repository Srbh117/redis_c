package main

import (
	"fmt"
	"net"
	"os"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			os.Exit(1)
		}
		_, err = conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			os.Exit(1)
		}
	}
}
