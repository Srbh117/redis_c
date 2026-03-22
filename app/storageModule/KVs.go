package main

import (
	"fmt"
	"srbh117/myRedis_c/app/parser"
	"strconv"
	"sync"
	"time"
)

type Entry struct {
	value  string
	expiry time.Time
}

type KeyValue struct {
	store   map[string]Entry
	myLocks sync.RWMutex
}

var KV KeyValue = KeyValue{
	store: make(map[string]Entry),
}

func SET(KeyValString string, TimeToExpireString string) error {
	KeyValList, err := parser.MyParser(KeyValString)
	if err != nil {
		return fmt.Errorf("ERROR PARSING KEY STRING LIST: ", err)
	}

	TimeToExpireList, err := parser.MyParser(TimeToExpireString)
	if err != nil {
		return fmt.Errorf("ERROR PARSING TIME STRING :", err)
	}

	if KeyValList[0] != "SET" {
		return fmt.Errorf("SET FUNCTIONALITY NOT REQUIRED | WRONG INPUT PROVIDED PLEASE CHECK STRING", KeyValList)
	}

	if len(KeyValList)-1 != 2*(len(TimeToExpireList)-1) {
		return fmt.Errorf("Mismatch in Data. Not every KV comes up with an Expiry. Please Check")
	}

	fmt.Println(KeyValList, TimeToExpireList)
	i := 1
	j := 1

	func() {
		KV.myLocks.Lock()
		defer KV.myLocks.Unlock()
		for i < len(KeyValList) && j < len(TimeToExpireList) {
			timeVal, err := strconv.Atoi(TimeToExpireList[j])
			if err != nil {
				return
			}

			KV.store[KeyValList[i]] = Entry{KeyValList[i+1], time.Now().Add(time.Millisecond * time.Duration(timeVal))}

			KV.store[KeyValList[i]] = Entry{KeyValList[i+1], time.Now().Add(time.Millisecond * time.Duration(timeVal))}

			i += 2
			j++
		}
	}()
	fmt.Println(KV.store)
	return nil

}

func GET(key string) (error, string) {
	v, ok := KV.store[key]
	if ok == false {
		return fmt.Errorf("KEY DOES NOT EXIST"), ""
	}
	if time.Now().After(v.expiry) == true {
		return fmt.Errorf("TIME ELAPSED, KEY EXPIRED"), ""
	}

	return nil, (KV.store[key].value)
}

func main() {
	var KV_STRING string = "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n"
	var TIME_STRING string = "*2\r\n$2\r\nPX\r\n$4\r\n1000\r\n"
	err := SET(KV_STRING, TIME_STRING)
	if err != nil {
		fmt.Print(err)
	}

	time.Sleep(5 * time.Second)
	err, val := GET("key")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(val)

}
