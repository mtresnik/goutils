package goutils

import (
	"fmt"
	"testing"
)

func TestDeepFlatten(t *testing.T) {
	var tempArray = []any{[]any{[]any{0}}, 1, 2, 3, []any{[]any{[]any{4}}}, []any{5, 6, 7}}
	fmt.Println(tempArray)
	var flattened = DeepFlatten(tempArray)
	fmt.Println(flattened)
}

func TestFloat64ArrayToString(t *testing.T) {
	tempArray := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println(Float64ArrayToString(tempArray, "<", ">"))
}

func TestIntArrayToString(t *testing.T) {
	tempArray := []int{1, 2, 3}
	fmt.Println(ArrayToString(tempArray))
}
