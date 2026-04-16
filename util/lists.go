package util

var store map[string][]string = make(map[string][]string)

func RPUSH(key string, val []string) int {
	store[key] = append(store[key], val...)
	return len(store[key])
}

func LRANGE(key string, start, stop int) string {
	val, ok := store[key]
	if ok != true {
		return ParseList([]string{})
	}
	if start >= len(val) {
		return ParseList([]string{})
	}
	if stop >= len(val) {
		stop = len(val) - 1
	}
	return ParseList(val[start : stop+1])
}
