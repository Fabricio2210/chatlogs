package changedateformat

import (
	"errors"
	"time"
)

func ChangeDateFormat(dateString string) (string, error) {
	parsedDateString := dateString[:4] + "-" + dateString[5:7] + "-" + dateString[8:]
	date, err := time.Parse("2006-01-02", parsedDateString)
	if err != nil {
		return dateString, errors.New("format date is invalid")
	}
	rfc3339Date := date.Format(time.RFC3339)
	return rfc3339Date, nil
}
