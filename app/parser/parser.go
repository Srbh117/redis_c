package parser

import (
	"fmt"
	"strings"
)

var SEPERATOR string = "\r\n"

func MyParser(userInput string) ([]string, error) {
	if len(userInput) == 0 {
		return []string{}, fmt.Errorf("NO STRING PROVIDED")
	}
	var LEN int = int(userInput[1]) - int('0')
	userInputSlice := make([]string, 0, LEN)
	curr_idx := 3 + len(SEPERATOR)
	for LEN > 0 {
		curr_string_len := int(userInput[curr_idx]) - int('0')
		curr_idx += len(SEPERATOR) + 1
		next_idx := strings.Index(userInput[curr_idx:], SEPERATOR)

		if next_idx != curr_string_len {
			return []string{}, fmt.Errorf("BAD STRING FORMATTING, LENGTH OF STRING IS WRONG")
		}
		userInputSlice = append(userInputSlice, userInput[curr_idx:curr_idx+next_idx])
		curr_idx = curr_idx + next_idx + len(SEPERATOR) + 1
		LEN -= 1
	}
	return userInputSlice, nil

}
