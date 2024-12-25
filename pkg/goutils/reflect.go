package goutils

import (
	"fmt"
	"reflect"
	"strconv"
)

func GetElemType(arr interface{}) reflect.Type {
	return reflect.TypeOf(arr).Elem()
}

func ToInt(a any) (int, error) {
	if val, ok := a.(int); ok {
		return val, nil
	}
	if s, ok := a.(string); ok {
		if s == "" {
			return 0, fmt.Errorf("string is empty")
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("cannot convert string to int")
		}
		return i, nil
	}
	return 0, fmt.Errorf("cannot convert to int")
}

func ToIntOrDefault(a any, def int) int {
	ret, err := ToInt(a)
	if err != nil {
		return def
	}
	return ret
}

func ToFloat64(a any) (float64, error) {
	if val, ok := a.(float64); ok {
		return val, nil
	}
	if s, ok := a.(string); ok {
		if s == "" {
			return 0, fmt.Errorf("string is empty")
		}
		i, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, fmt.Errorf("cannot convert string to float64")
		}
		return i, nil
	}
	return 0, fmt.Errorf("cannot convert to float64")
}

func ToFloat64OrDefault(a any, def float64) float64 {
	f, err := ToFloat64(a)
	if err != nil {
		return def
	}
	return f
}
