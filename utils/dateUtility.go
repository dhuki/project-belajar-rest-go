package utils

import "time"

const FORMAT_DATE_TIME = "02/01/2006 15:04:05"
const FORMAT_DATE = "02/01/2006"

func ConvertStringtoDate(format string, date string) (time.Time, error) {
	t, err := time.Parse(format, date)
	if err != nil {
		return time.Now(), err
	}

	return t, nil
}
