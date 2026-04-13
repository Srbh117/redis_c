package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var _ = net.Listen
var _ = os.Exit

var SEP = "\r\n"
var NULL_STRING = "$-1\r\n"

var KEY_EXPIRED = errors.New("Key expired")

/*
	Store ->

	Saves all keys. Per KEY -> VALUE AND AN EXPIRY (THIS GOOD). IF EXPIRY IS NIL THEN INFI, ELSE TILL EXPIRY.
	WHILE FETCH IF EXPIRY > TIME.NOW() THEN DELETE KEY.

*/

type KV struct {
	store map[string]Entry
}

type Entry struct {
	value  string
	expiry *time.Time
}

func NewKV() KV {
	return KV{
		store: make(map[string]Entry),
	}
}

var myStore KV = NewKV()

func Set(parsedResp []string) error {
	log.Println("Argument:", parsedResp)
	if len(parsedResp) < 3 {
		return fmt.Errorf("LENGTH OF PARSED RESP LESS THAN 3. NO KEY EXISTS")
	}

	if len(parsedResp[1]) == 0 || len(parsedResp[2]) == 0 {
		return fmt.Errorf("LENGTH OF Provided KEY | VALUE IS 0")
	}

	var timeToExpire *time.Time

	if len(parsedResp) <= 3 {
		timeToExpire = nil
	}
	if parsedResp[3] != "EX" || parsedResp[3] != "PX" {
		timeToExpire = nil
	}

	timeDur, err := strconv.Atoi(parsedResp[4])
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	var dur time.Duration
	if parsedResp[3] == "EX" {
		dur = time.Second
	} else {
		dur = time.Millisecond
	}
	t := time.Now().Add(time.Duration(timeDur) * dur)
	timeToExpire = &t
	_, ok := myStore.store[parsedResp[1]]
	if ok == false {
		myStore.store[parsedResp[1]] = Entry{
			value:  parsedResp[2],
			expiry: timeToExpire,
		}
	}
	myStore.store[parsedResp[1]] = Entry{value: parsedResp[2], expiry: timeToExpire}

	return nil
}

func GET(key string) (any, error) {
	if len(key) == 0 {
		return "", fmt.Errorf("EMPTY KEY SEND")
	}
	v, ok := myStore.store[key]
	if ok == false {
		return "", fmt.Errorf("KEY DOESN't EXIST")
	}
	if v.expiry == nil {
		return v.value, nil
	}

	if v.expiry.After(time.Now()) == false {
		delete(myStore.store, key)
		return "", KEY_EXPIRED
	} else {
		return v.value, nil
	}

}

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
			conn.Write([]byte("+PONG\r\n"))
		}

		if receiver_arr[0] == "GET" {
			val, err := GET(receiver_arr[1])
			if err != nil {
				conn.Write([]byte("$-1\r\n"))
			}

			if err == KEY_EXPIRED {
				conn.Write([]byte("$-1\r\n"))
			}
			stringVal, ok := val.(string)
			if ok != true {

				conn.Write([]byte(NULL_STRING))
			}
			conn.Write([]byte(ConvertSingleString(stringVal)))
		}

		if receiver_arr[0] == "SET" {
			err := Set(receiver_arr)
			if err != nil {
				conn.Write([]byte("$-1\r\n"))
			}
			conn.Write([]byte("+OK\r\n"))
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
