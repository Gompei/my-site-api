package pkg

import "time"

const timeFormat = "2006-01-02"

func CreateTimeStamp() string {
	return format(time.Now())
}

func format(t time.Time) string {
	return t.Format(timeFormat)
}
