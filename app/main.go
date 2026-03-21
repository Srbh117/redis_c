package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

var SEPERATOR string = "\r\n"

func MyParser(userInput string) ([]string, error) {
	if len(userInput) == 0 {
		return []string{}, fmt.Errorf("NO STRING PROVIDED")
	}
	var LEN int = int(userInput[1]) - int('0')
	userInputSlice := make([]string, 0, LEN)
	curr_idx := 3 + len(SEPERATOR)
	for LEN > 0 {
		curr_string_len := int(userInput[curr_idx]) - int('0')
		curr_idx += len(SEPERATOR) + 1
		next_idx := strings.Index(userInput[curr_idx:], SEPERATOR)

		if next_idx != curr_string_len {
			return []string{}, fmt.Errorf("BAD STRING FORMATTING, LENGTH OF STRING IS WRONG")
		}
		userInputSlice = append(userInputSlice, userInput[curr_idx:curr_idx+next_idx])
		curr_idx = curr_idx + next_idx + len(SEPERATOR) + 1
		LEN -= 1
	}
	return userInputSlice, nil

}

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
	if string((*buff)[:n]) == "PING\n" {
		conn.Write([]byte("+PONG\r\n"))
		return
	}
	userInputString := string((*buff)[:n])
	userInputSlice, err := MyParser(userInputString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("USERINPUTSLICE:\t", userInputSlice)
	conn.Write([]byte(userInputSlice[1]))

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
