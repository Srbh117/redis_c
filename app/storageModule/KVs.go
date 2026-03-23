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
	store   map[string]NewEntry
	myLocks sync.RWMutex
}

var KV KeyValue = KeyValue{
	store: make(map[string]NewEntry),
}

func LPush(key string, insertionString string) error {

	KV.myLocks.Lock()
	defer KV.myLocks.Unlock()
	_, ok := KV.store[key]
	if !ok {
		KV.store[key] = NewEntry{value: []string{insertionString}}
		return nil
	}
	switch val := KV.store[key].value.(type) {
	case []string:
		arr := []string{insertionString}
		arr = append(arr, val...)
		KV.store[key] = NewEntry{value: arr}
		return nil

	default:
		return fmt.Errorf("Type of Value is not []string ")
	}

}

func RPush(key string, insertionString string) error {

	KV.myLocks.Lock()
	defer KV.myLocks.Unlock()
	_, ok := KV.store[key]
	if ok != true {
		KV.store[key] = NewEntry{
			value: []string{insertionString},
		}
		return nil
	}
	switch val := KV.store[key].value.(type) {
	case []string:
		val = append(val, insertionString)
		KV.store[key] = NewEntry{value: val}
	default:
		return fmt.Errorf("Type of Value for Key Not Equal to []string")
	}
	return nil
}

// Command	What it does
// LPUSH key val	Push to head
// RPUSH key val	Push to tail
// LPOP key	Remove + return from head
// RPOP key	Remove + return from tail
// // LRANGE key start stop	Get elements start..stop (0-indexed, -1 = last)
// LLEN key

func LPOP(key string) (string, error) {

	KV.myLocks.Lock()
	defer KV.myLocks.Unlock()
	_, ok := KV.store[key]
	if !ok {
		KV.store[key] = NewEntry{value: []string{}}
		return "", nil
	}

	switch val := KV.store[key].value.(type) {
	case []string:
		if len(val) == 0 {
			KV.store[key] = NewEntry{value: []string{}}
			return "", nil
		}
		f := val[0]
		KV.store[key] = NewEntry{value: val[1:]}
		return f, nil
	default:
		return "", fmt.Errorf("Type is NOT []string")
	}

}

func RPOP(key string) (string, error) {

	KV.myLocks.Lock()
	defer KV.myLocks.Unlock()
	_, ok := KV.store[key]
	if !ok {
		KV.store[key] = NewEntry{value: []string{}}
		return "", nil
	}

	switch val := KV.store[key].value.(type) {
	case []string:
		if len(val) == 0 {
			KV.store[key] = NewEntry{value: []string{}}
			return "", nil
		}
		f := val[len(val)-1]
		KV.store[key] = NewEntry{value: val[:len(val)-1]}
		return f, nil
	default:
		return "", fmt.Errorf("Type is NOT []string")
	}

}

func Set(KeyValList []string) error {
	if len(KeyValList) != 5 && len(KeyValList) != 3 {
		return fmt.Errorf("KEY VAL LIST LENGTH ISN't 5 OR 3, SOMETHING IS MISSING")
	}

	func() {
		KV.myLocks.Lock()
		defer KV.myLocks.Unlock()
		if len(KeyValList) == 3 {
			KV.store[KeyValList[1]] = NewEntry{value: KeyValList[2]}
			return
		}

		var timeToExpire time.Time
		timeVal, err := strconv.Atoi(KeyValList[4])
		if err != nil {
			return
		}
		if KeyValList[3] == "PX" {
			timeToExpire = time.Now().Add(time.Millisecond * time.Duration(timeVal))
		} else {
			timeToExpire = time.Now().Add(time.Second * time.Duration(timeVal))
		}

		KV.store[KeyValList[1]] = NewEntry{
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

	if KV.store[key].expiry.IsZero() || time.Now().After(v.expiry) == false {
		switch val := KV.store[key].value.(type) {
		case string:
			return nil, val
		default:
			return fmt.Errorf("WRONGTYPE Operation against a key holding the wrong kind of value"), ""

		}
	}

	return fmt.Errorf(" Key Expired "), ""
}
