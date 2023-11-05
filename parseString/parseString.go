package parsestring

import (
	"errors"
	"strings"
)

type ParsedInfo struct {
	Date    string
	Hour    string
	Name    string
	Message string
}

func Parsestring(str string) (ParsedInfo, error) {
	substrings := strings.SplitN(str, "|", 2)
	if len(substrings) != 2 {
		return ParsedInfo{}, errors.New("input string format is invalid")
	}
	substringDateAndHourUnparsed := substrings[0]
	substringNameAndMessageUnparsed := substrings[1]

	substringDateAndHourParsed := strings.Split(substringDateAndHourUnparsed, " ")
	if len(substringDateAndHourParsed) != 3 {
		return ParsedInfo{}, errors.New("invalid date and hour format")
	}
	substringNameAndMessageParsed := strings.SplitN(substringNameAndMessageUnparsed, ":", 2)
	if len(substringNameAndMessageParsed) != 2 {
		return ParsedInfo{}, errors.New("invalid name and message format")
	}
	substringDate := substringDateAndHourParsed[0]
	substringHour := substringDateAndHourParsed[1]
	substringName := substringNameAndMessageParsed[0]
	substringMessage := substringNameAndMessageParsed[1]
	parsedInfo := ParsedInfo{
		Date:    substringDate,
		Hour:    substringHour,
		Name:    substringName,
		Message: substringMessage,
	}
	return parsedInfo, nil
}
