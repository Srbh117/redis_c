package main

import (
	"fmt"
	"io"
	"net"
	"srbh117/myRedis_c/app/parser"
	storagemodule "srbh117/myRedis_c/app/storageModule"
	"strings"
)

var SEPERATOR string = "\r\n"

func handleTCP(conn net.Conn, curr_idx *int, buff *[]byte) {
	defer conn.Close()
	fmt.Println("NET CALL MADE")
	n, err := conn.Read(*buff)
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Errorf("Error Reading Conn Buff %v", err)
	}

	userInputString := strings.ReplaceAll(string((*buff)[:n]), "\r\n", "\n")
	userInputString = strings.ReplaceAll(userInputString, "\n", "\r\n")

	userInputSlice, err := parser.MyEasyParser(userInputString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("USERINPUTSLICE:\t", userInputSlice)

	if userInputSlice[0] == "ECHO" {
		conn.Write([]byte(userInputSlice[1]))
	}
	if userInputSlice[0] == "PING" {
		conn.Write([]byte("+PONG"))
	}
	if userInputSlice[0] == "SET" {
		storagemodule.Set(userInputSlice)
	}
	if userInputSlice[0] == "GET" {
		err, str := storagemodule.GET(userInputSlice[1])
		if err != nil {
			conn.Write([]byte(err.Error()))
		} else {
			conn.Write([]byte(str))
		}
	}

}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:42069")
	if err != nil {
		return
	}
	curr_idx := 0
	buff := make([]byte, 1024)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Errorf("ERROR BINDING TO PORT 42069")
		}

		go handleTCP(conn, &curr_idx, &buff)

	}
}
