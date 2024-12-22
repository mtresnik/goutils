package goutils

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func GetOrDefault[T any](a []T, index int, defaultValue T) T {
	if index < len(a) {
		return a[index]
	}
	return defaultValue
}

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

func HashFloats(values ...float64) int64 {
	hasher := sha256.New()

	for _, value := range values {
		bits := math.Float64bits(value)
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, bits)
		_, _ = hasher.Write(buf)
	}

	size := len(values)
	sizeBuf := make([]byte, 8)
	binary.BigEndian.PutUint64(sizeBuf, uint64(size))
	_, _ = hasher.Write(sizeBuf)

	fullHash := hasher.Sum(nil)
	hash1 := binary.BigEndian.Uint64(fullHash[:8])
	hash2 := binary.BigEndian.Uint64(fullHash[8:16])

	finalHash := int64(hash1 ^ hash2)
	return finalHash
}

func MapNotNil[T any, U any](arr []T, f func(T) *U) []*U {
	var result = make([]*U, 0)
	for _, v := range arr {
		u := f(v)
		if u != nil {
			result = append(result, u)
		}
	}
	return result
}

func Map[T any, U any](arr []T, f func(T) U) []U {
	var result = make([]U, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}

func Filter[T any](arr []T, f func(T) bool) []T {
	var result = make([]T, 0)
	for _, v := range arr {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func MinBy[T any](arr []T, f func(T) float64) T {
	minResult := arr[0]
	for _, v := range arr {
		if f(v) < f(minResult) {
			minResult = v
		}
	}
	return minResult
}

func MaxBy[T any](arr []T, f func(T) float64) T {
	maxResult := arr[0]
	for _, v := range arr {
		if f(v) > f(maxResult) {
			maxResult = v
		}
	}
	return maxResult
}

func SumBy[T any](arr []T, f func(T) float64) float64 {
	sum := 0.0
	for _, v := range arr {
		sum += f(v)
	}
	return sum
}

func IndexOf[T any](arr []T, f func(T) bool) int {
	for i, v := range arr {
		if f(v) {
			return i
		}
	}
	return -1
}

func Contains[T any](arr []T, f func(T) bool) bool {
	return IndexOf(arr, f) != -1
}

func Unique[T comparable](arr []T) []T {
	var result = make([]T, 0)
	for _, v := range arr {
		if !Contains(result, func(t T) bool { return t == v }) {
			result = append(result, v)
		}
	}
	return result
}

func All[T any](arr []T, f func(T) bool) bool {
	for _, v := range arr {
		if !f(v) {
			return false
		}
	}
	return true
}

func Any[T any](arr []T, f func(T) bool) bool {
	for _, v := range arr {
		if f(v) {
			return true
		}
	}
	return false
}

func Chunk[T any](slice []T, chunkSize int) [][]T {
	if chunkSize <= 0 {
		return nil
	}
	chunks := make([][]T, 0, (len(slice)+chunkSize-1)/chunkSize)
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func Reduce[T, U any](slice []T, initial U, f func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = f(result, v)
	}
	return result
}

func Partition[T any](slice []T, predicate func(T) bool) ([]T, []T) {
	var trueItems, falseItems []T
	for _, item := range slice {
		if predicate(item) {
			trueItems = append(trueItems, item)
		} else {
			falseItems = append(falseItems, item)
		}
	}
	return trueItems, falseItems
}

func Indices[T any](slice []T) chan int {
	c := make(chan int)
	go func() {
		for i, _ := range slice {
			c <- i
		}
		close(c)
	}()
	return c
}

func RangeOfIntsByStep(start int, end int, step int) []int {
	var result = make([]int, 0)
	for i := start; i < end; i += step {
		result = append(result, i)
	}
	return result
}

func RangeOfInts(start int, end int) []int {
	return RangeOfIntsByStep(start, end, 1)
}
