package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

var _ = net.Listen
var _ = os.Exit

var SEP = "\n"

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		return
	}

	conn, err := l.Accept()
	if err != nil {
		os.Exit(1)
	}

	buff := make([]byte, 1024)
	conn.Read(buff)
	curr_idx := 0
	for bytes.Index(buff[curr_idx:], []byte(SEP)) != -1 {
		_, err = conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			os.Exit(1)
		}
		curr_idx = bytes.Index(buff[curr_idx:], []byte(SEP))
	}

}
