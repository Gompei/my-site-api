package object

import "time"

type TimeStamp struct {
	timeStamp string
}

func NewTimeStamp() *TimeStamp {
	return &TimeStamp{
		timeStamp: format(time.Now()),
	}
}

func format(t time.Time) string {
	return t.Format(time.RFC3339)
}
