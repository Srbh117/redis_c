package util

var store map[string][]string

func RPUSH(key string, val string) int {
	store[key] = append(store[key], val)
	return len(store[key])
}
