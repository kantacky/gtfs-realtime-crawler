package lib

import "time"

func ISO8601(datetime time.Time) string {
	return datetime.Format("2006-01-02T15:04:05-07:00")
}

func ParseISO8601(datetimeString string) *time.Time {
	datetime, err := time.Parse("2006-01-02T15:04:05-07:00", datetimeString)
	if err != nil {
		return nil
	}
	return &datetime
}
