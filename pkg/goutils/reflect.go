package goutils

import "reflect"

func GetElemType(arr interface{}) reflect.Type {
	return reflect.TypeOf(arr).Elem()
}
