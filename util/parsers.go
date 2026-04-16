package util

import "strconv"

func ParseInt(val int) string {
	str := strconv.Itoa(val)
	return str + "\r\n"
}
