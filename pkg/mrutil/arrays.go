package mrutil

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ArrayToString(a interface{}) string {
	t := GetElemType(a)
	switch t.String() {
	case "int":
		return IntArrayToString(a.([]int))
	case "int64":
		return Int64ArrayToString(a.([]int64))
	case "float64":
		return Float64ArrayToString(a.([]float64))
	case "string":
		return fmt.Sprintf("%v", a.([]string))
	}
	return fmt.Sprint(a)
}

func IntArrayToString(a []int, braces ...string) string {
	retArray := make([]string, len(a))
	for i, v := range a {
		retArray[i] = strconv.Itoa(v)
	}
	if len(braces) > 1 {
		return fmt.Sprintf("%s%v%s", braces[0], strings.Join(retArray, ", "), braces[1])
	}
	return fmt.Sprintf("[%v]", strings.Join(retArray, ", "))
}

func Int64ArrayToString(a []int64, braces ...string) string {
	retArray := make([]string, len(a))
	for i, v := range a {
		retArray[i] = strconv.FormatInt(v, 10)
	}
	if len(braces) > 1 {
		return fmt.Sprintf("%s%v%s", braces[0], strings.Join(retArray, ", "), braces[1])
	}
	return fmt.Sprintf("[%v]", strings.Join(retArray, ", "))
}

func Float64ArrayToString(a []float64, braces ...string) string {
	retArray := make([]string, len(a))
	for i, v := range a {
		retArray[i] = strconv.FormatFloat(v, 'f', -1, 64)
	}
	if len(braces) > 1 {
		return fmt.Sprintf("%s%v%s", braces[0], strings.Join(retArray, ", "), braces[1])
	}
	return fmt.Sprintf("[%v]", strings.Join(retArray, ", "))
}

func UnpackArray(s any) []any {
	v := reflect.ValueOf(s)
	r := make([]any, v.Len())
	for i := 0; i < v.Len(); i++ {
		r[i] = v.Index(i).Interface()
	}
	return r
}

func HasArrays(root []any) bool {
	for i := 0; i < len(root); i++ {
		if reflect.TypeOf(root[i]) == reflect.TypeOf(root) {
			return true
		}
	}
	return false
}

func Flatten(root []any) []any {
	var count = 0
	for i := 0; i < len(root); i++ {
		if reflect.TypeOf(root[i]) == reflect.TypeOf(root) {
			childArray := UnpackArray(root[i])
			count += len(childArray)
		} else {
			count++
		}
	}
	var retArray = make([]any, count)
	var outerIndex = 0
	for i := 0; i < len(root); i++ {
		if reflect.TypeOf(root[i]) == reflect.TypeOf(root) {
			childArray := UnpackArray(root[i])
			for j := 0; j < len(childArray); j++ {
				retArray[outerIndex] = childArray[j]
				outerIndex++
			}
		} else {
			retArray[outerIndex] = root[i]
			outerIndex++
		}
	}
	return retArray
}

func DeepFlatten(root []any) []any {
	var retArray = root
	for HasArrays(retArray) {
		retArray = Flatten(retArray)
	}
	return retArray
}
