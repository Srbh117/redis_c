package storagemodule

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Entry struct {
	value  string
	expiry time.Time
}

type NewEntry struct {
	value  interface{}
	expiry time.Time
}

type KeyValue struct {
	store   map[string]Entry
	myLocks sync.RWMutex
}

var KV KeyValue = KeyValue{
	store: make(map[string]Entry),
}

func (Entry *NewEntry) LPush(value string) {
	existingList, _ := Entry.value.([]string)

	arr := []string{value}
	arr = append(arr, existingList...)
	Entry.value = arr
}

func (Entry *NewEntry) RPush(value string) {
	existingList, ok := Entry.value.([]string)
	if ok != true {
		Entry.value = []string{}
	}

	existingList = append(existingList, value)
	Entry.value = existingList
}

func Set(KeyValList []string) error {
	if len(KeyValList) != 5 && len(KeyValList) != 3 {
		return fmt.Errorf("KEY VAL LIST LENGTH ISN't 5 OR 3, SOMETHING IS MISSING")
	}

	func() {
		KV.myLocks.Lock()
		defer KV.myLocks.Unlock()
		var timeToExpire time.Time
		timeVal, err := strconv.Atoi(KeyValList[4])
		if err != nil {
			return
		}
		if len(KeyValList) == 3 {
			KV.store[KeyValList[1]] = Entry{value: KeyValList[2]}
		}

		if KeyValList[3] == "PX" {
			timeToExpire = time.Now().Add(time.Millisecond * time.Duration(timeVal))
		} else {
			timeToExpire = time.Now().Add(time.Second * time.Duration(timeVal))
		}

		KV.store[KeyValList[1]] = Entry{
			value:  KeyValList[2],
			expiry: timeToExpire,
		}
	}()
	return nil
}

func GET(key string) (error, string) {
	v, ok := KV.store[key]
	if ok == false {
		return fmt.Errorf("KEY DOES NOT EXIST"), ""
	}

	if KV.store[key].expiry.IsZero() {
		return nil, KV.store[key].value
	}
	if time.Now().After(v.expiry) == true {
		return fmt.Errorf("TIME ELAPSED, KEY EXPIRED"), ""
	}

	return nil, (KV.store[key].value)
}
