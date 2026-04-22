package util

import "slices"

var store map[string][]string = make(map[string][]string)

func RPUSH(key string, val []string) int {
	store[key] = append(store[key], val...)
	return len(store[key])
}

func BLPOP(key string, time int) string {
	if time != 0 {
		return "ASS"
	} else {
	loop:
		for {
			_, ok := store[key]
			if ok == false {
				goto loop
			}
			if ok == true && len(store[key]) == 0 {
				goto loop
			}
			val := LPOP(key)
			return ParseList([]string{key, val})
		}
	}

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

func LLEN(key string) int {
	val, ok := store[key]
	if ok == false {
		return 0
	}
	return len(val)
}

func LPOP(key string) string {
	v, ok := store[key]
	if ok == false || len(v) == 0 {
		return ""
	}

	f := v[0]
	store[key] = store[key][1:]
	return f
}

func LPOP_Ranged(key string, rangeVal int) []string {
	v, ok := store[key]
	if ok == false || len(v) == 0 {
		return []string{}
	}
	if rangeVal > len(store[key]) {
		f := store[key]
		store[key] = []string{}
		return f
	}
	f := store[key][:rangeVal]
	store[key] = store[key][rangeVal:]
	return f
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
