package util

import "slices"

var store map[string][]string = make(map[string][]string)

func RPUSH(key string, val []string) int {
	store[key] = append(store[key], val...)
	return len(store[key])
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func LPUSH(key string, val []string) int {
	for _, v := range val {
		store[key] = slices.Insert(store[key], 0, v)
	}
	return len(store[key])
}

func LRANGE(key string, start, stop int) string {
	val, ok := store[key]
	if ok != true {
		return ReturnEmptyString()
	}

	if start < 0 {
		if abs(start) > len(store[key]) {
			start = 0
		} else {
			start += len(store[key])
		}
	}
	if stop < 0 {
		if abs(stop) > len(store[key]) {
			stop = 0
		} else {
			stop += len(store[key])
		}
	}

	if start >= len(val) {
		return ReturnEmptyString()
	}
	if stop >= len(val) {
		stop = len(val) - 1
	}
	return ParseList(val[start : stop+1])
}
