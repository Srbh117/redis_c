package main

import (
	"fmt"
	"srbh117/myRedis_c/app/parser"
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

	//SizeChecksfor KeyVal and TimeBound
	//Since KeyVal-1==2*(TimeBound-1)

	if len(KeyValList)-1 != 2*(len(TimeToExpireList)-1) {
		return fmt.Errorf("Mismatch in Data. Not every KV comes up with an Expiry. Please Check")
	}
	i := 1
	j := 1
	for i < len(KeyValList) && j < len(TimeToExpireList) {

	}

	return nil

}

func main() {
	var KV_STRING string = "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n"
	var TIME_STRING string = "*2\r\n$2\r\nPX\r\n$4\r\n1000\r\n"
	err := SET(KV_STRING, TIME_STRING)
	if err != nil {
		fmt.Print(err)
	}

}
