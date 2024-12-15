package mrutil

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func StringStartsWith(one string, other string) bool {
	return strings.Index(one, other) == 0
}

func StringEndsWith(one string, other string) bool {
	lastIndex := utf8.RuneCountInString(one)
	return strings.LastIndex(one, other) == lastIndex-utf8.RuneCountInString(other)
}

func SubstringToEnd(input string, startIndex int) string {
	return Substring(input, startIndex, utf8.RuneCountInString(input))
}

func Substring(input string, startIndex int, endIndex int) string {
	runeSlice := []rune(input)
	if startIndex >= len(runeSlice) {
		return ""
	}
	var length = endIndex - startIndex
	if endIndex > len(runeSlice) {
		length = len(runeSlice) - startIndex
		endIndex = startIndex + length
	}
	return string(runeSlice[startIndex:endIndex])
}

func FindRemainingStrings(test string, key string) []string {
	if utf8.RuneCountInString(key) == 0 {
		return make([]string, 0)
	}
	if strings.Contains(test, key) == false {
		return make([]string, 0)
	}
	if utf8.RuneCountInString(test) == utf8.RuneCountInString(key) {
		return make([]string, 0)
	}
	var index = strings.Index(test, key)
	if index == 0 {
		return []string{SubstringToEnd(test, utf8.RuneCountInString(key))}
	}
	if StringEndsWith(test, key) {
		return []string{Substring(test, 0, utf8.RuneCountInString(test)-utf8.RuneCountInString(key))}
	} else {
		return []string{
			Substring(test, 0, index),
			SubstringToEnd(test, index+utf8.RuneCountInString(key)),
		}
	}
}

func SmartComplexString(c complex128) string {
	if imag(c) == 0 {
		return strconv.FormatFloat(real(c), 'G', 5, 64)
	}
	return strconv.FormatComplex(c, 'G', 5, 64)
}

func SliceToString(array []int64) string {
	retString := "["
	for i, v := range array {
		retString += strconv.Itoa(int(v))
		if i < len(array)-1 {
			retString += ", "
		}
	}
	retString += "]"
	return retString
}
