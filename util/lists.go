package util

type LIST struct {
	list []any
}

var l LIST

func RPUSH(val string) int {
	l.list = append(l.list, val)
	return len(l.list)
}
