package date

import "time"

func GetNowAfter() time.Time {
	now := time.Now()
	location := time.Local
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, location)
	return endTime
}
