package mrutil

import "reflect"

func GetElemType(arr interface{}) reflect.Type {
	return reflect.TypeOf(arr).Elem()
}
