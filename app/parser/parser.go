package parser

import (
	"fmt"
	"strconv"
	"strings"
)

var SEPERATOR string = "\r\n"

func MyParser(userInput string) ([]string, error) {
	fmt.Printf(userInput)
	if len(userInput) == 0 {
		return []string{}, fmt.Errorf("NO STRING PROVIDED")
	}
	var LEN int = int(userInput[1]) - int('0')
	userInputSlice := make([]string, 0, LEN)
	curr_idx := 3 + len(SEPERATOR)
	for LEN > 0 {
		curr_string_len := int(userInput[curr_idx] - '0')
		if curr_string_len == 0 {
			return []string{}, fmt.Errorf("I DO NOT WANT EMPTY STRINGS!!!!!!!!")
		}
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

func ParseNumbers(numberInput string) (int, error) {
	if len(numberInput) == 0 {
		return -1, fmt.Errorf("Zero Sized Number Sent")
	}
	if numberInput[0] != '*' && numberInput[0] != '$' {
		return -1, fmt.Errorf("Incorrect Formatting Please Check")
	}

	number, err := strconv.Atoi(numberInput[1:])
	if err != nil {
		return -1, err
	}
	return number, nil
}

func MyEasyParser(userInput string) ([]string, error) {
	ListOfStrings := strings.Split(userInput, SEPERATOR)
	NumberOfElement, err := ParseNumbers(ListOfStrings[0])
	if err != nil {
		return []string{}, err
	}
	var ans []string
	i := 1
	for i <= 2*NumberOfElement {
		next_len, err := ParseNumbers(ListOfStrings[i])
		if err != nil {
			return []string{}, err
		}
		i += 1
		if len(ListOfStrings[i]) != next_len {
			return []string{}, fmt.Errorf("String Length Mismatch")
		}
		ans = append(ans, ListOfStrings[i])
		i += 1
	}
	return ans, nil
}
