package goutils

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

func TestStringEndsWith(t *testing.T) {
	testString := "abc123"
	fmt.Println(StringEndsWith(testString, "123"))
}

func TestSubstring(t *testing.T) {
	testString := "abc123"
	endIndex := strings.Index(testString, "123")
	fmt.Println(Substring(testString, 0, endIndex))
}

func TestDelete(t *testing.T) {
	var slice = []int{0, 1, 2, 3}
	slice = slices.Delete(slice, 2, 3)
	fmt.Println(slice)
}
