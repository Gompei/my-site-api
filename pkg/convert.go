package pkg

import (
	"encoding/json"
	"reflect"
)

func InterfaceToJson(i interface{}) (string, error) {
	j, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(j), err
}

func InterfaceToStruct(str string, i interface{}) (interface{}, error) {
	err := json.Unmarshal([]byte(str), i)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func MapToStruct(m map[string]interface{}, i interface{}) interface{} {
	v := reflect.Indirect(reflect.ValueOf(i))
	for n, e := range m {
		v.FieldByName(n).Set(reflect.ValueOf(e))
	}
	return i
}

//func MapToStruct(m map[string]interface{}, i interface{}) (interface{}, error) {
//	j, err := json.Marshal(m)
//	if err != nil {
//		return nil, err
//	}
//	err = json.Unmarshal(j, i)
//	if err != nil {
//		return nil, err
//	}
//	return i, err
//}