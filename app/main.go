package main

import (
	"fmt"
	"io"
	"net"
)

func handleTCP(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Errorf("Error Reading Conn Buff %v", err)
	}

	if string(buff[:n]) == "PING\n" {
		conn.Write([]byte("+PONG\r\n"))
	}

}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:42069")
	if err != nil {
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Errorf("ERROR BINDING TO PORT 42069")
		}

		go handleTCP(conn)

	}
}
