package pkg

import "time"

func CreateTimeStamp() string {
	return format(time.Now())
}

func format(t time.Time) string {
	return t.Format(time.RFC3339)
}
