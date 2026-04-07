package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

var _ = net.Listen
var _ = os.Exit

var SEP = "\r\n"

func parseString(resp string) []string {
	//*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n
	myArr := strings.Split(resp, SEP)
	ans := []string{}
	for i := 2; i < len(myArr); i += 2 {
		ans = append(ans, strings.TrimSpace(myArr[i]))
	}
	return ans

	//[ECHO,HEY]
}

func ConvertSingleString(word string) string {
	lenStr := "$" + strconv.Itoa(len(word)) + SEP + word + SEP
	return lenStr

}

func CreateString(resp []string) string {
	ans := ""
	lenArr := "*" + strconv.Itoa(len(resp))
	ans += lenArr
	ans += SEP
	for i := 0; i < len(resp); i++ {
		lenStr := "$" + strconv.Itoa(len(resp[i])) + SEP + resp[i] + SEP
		ans += lenStr
	}

	return ans
}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Errorf(err.Error())
			return
		}

		receiver_str := string(buff[:n])
		receiver_arr := parseString(receiver_str)
		if parseString(receiver_str)[0] == "ECHO" {
			conn.Write([]byte(ConvertSingleString(parseString(receiver_str)[1])))
		}
		if receiver_arr[0] == "PING" {
			conn.Write([]byte(ConvertSingleString("PONG")))
		}

	}
}

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
		go handleConnection(conn)
	}

	// arr := parseString("*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n")

	// for k, v := range arr {
	// 	fmt.Println(k, v, len(v))
	// }
	// str := CreateString(arr)
	// fmt.Println(str)

}
