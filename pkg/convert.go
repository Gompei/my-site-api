package pkg

import (
	"encoding/json"
	"strconv"
	"time"
)

func StringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func StringToTime(str string) time.Time {
	t, _ := time.Parse(timeFormat, str)
	return t
}

func InterfaceToJson(i interface{}) (string, error) {
	j, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(j), err
}
