package util

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

func LRANGE(key string, start, stop int) string {
	val, ok := store[key]
	if ok != true {
		return ParseList([]string{})
	}

	if start < 0 {
		if abs(start) > len(store[key]) {
			start = 0
		}
		start += len(store[key])
	}
	if stop < 0 {
		if abs(stop) > len(store[key]) {
			stop = 0
		}
		stop += len(store[key])
	}

	if start >= len(val) {
		return ParseList([]string{})
	}
	if stop >= len(val) {
		stop = len(val) - 1
	}
	return ParseList(val[start : stop+1])
}
