package util

import (
	"strconv"
)

var SEP = "\r\n"

func ParseInt(val int) string {
	str := strconv.Itoa(val)
	return ":" + str + "\r\n"
}

func ParseList(list []string) string {
	length := len(list)
	str := ""
	str += "*" + strconv.Itoa(length) + SEP
	for _, v := range list {
		curr_str := "$" + strconv.Itoa(len(v)) + SEP
		curr_str += v + SEP
		str += curr_str
	}
	return str
}

func ReturnEmptyString() string {
	return "*0" + SEP
}
