package utils

import (
	"encoding/json"
)

func UnSafeObjectToJsonString(obj interface{}) string {
	b, _ := json.Marshal(obj)
	return string(b)
}

func SafeObjectToJsonString(obj interface{}) (string, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func StringToObject[T interface{}](str string) (T, error) {
	var target T
	err := json.Unmarshal([]byte(str), &target)
	if err != nil {
		return target, err
	}
	return target, nil
}

func ObjectToObject(obj1 interface{}, obj2 interface{}) error {
	b, err := json.Marshal(obj1)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, obj2)
	if err != nil {
		return err
	}
	return nil
}
