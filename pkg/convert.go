package pkg

import (
	"encoding/json"
	"time"
)

func StringToTime(str string) time.Time {
	t, _ := time.Parse(time.RFC3339, str)
	return t
}

func InterfaceToJson(i interface{}) (string, error) {
	j, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(j), err
}
