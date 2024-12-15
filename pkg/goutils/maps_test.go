package goutils

import (
	"fmt"
	"testing"
)

func TestKeys(t *testing.T) {
	tempMap := map[string]int{}
	tempMap["a"] = 0
	tempMap["b"] = 1
	tempMap["c"] = 2
	tempMap["d"] = 3
	tempMap["e"] = 4
	tempMap["f"] = 5
	tempMap["g"] = 6
	keys := Keys(tempMap)
	fmt.Println(keys)
}

func TestMapToString(t *testing.T) {
	tempMap := map[string]int{}
	tempMap["a"] = 0
	tempMap["b"] = 1
	tempMap["c"] = 2
	tempMap["d"] = 3
	tempMap["e"] = 4
	tempMap["f"] = 5
	tempMap["g"] = 6
	fmt.Println(tempMap)
}
